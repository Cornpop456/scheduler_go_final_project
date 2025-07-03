package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func setRepeatWeekDays(weekDay *[8]bool, days []string) error {
	for _, s := range days {
		dayNum, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("invalid day number: %s", s)
		}
		if dayNum < 1 || dayNum > 7 {
			return fmt.Errorf("day of week must be between 1 and 7, got %d", dayNum)
		}

		weekDay[dayNum] = true
	}
	return nil
}

func setRepeatDays(day *[34]bool, days []string) error {
	for _, s := range days {
		dayNum, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("invalid day number: %s", s)
		}
		if dayNum < -2 || dayNum > 31 {
			return fmt.Errorf("day of month must be between -2 and 31, got %d", dayNum)
		}

		if dayNum < 0 {
			switch dayNum {
			case -2:
				day[predLastDayIndex] = true
			case -1:
				day[lastDayIndex] = true
			}

			continue
		}

		day[dayNum] = true
	}
	return nil
}

func setRepeatMonths(month *[13]bool, months []string) error {
	for _, s := range months {
		monthNum, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("invalid month number: %s", s)
		}
		if monthNum < 1 || monthNum > 12 {
			return fmt.Errorf("month must be between 1 and 12, got %d", monthNum)
		}

		month[monthNum] = true
	}
	return nil
}

func checkDate(t string, nextDate time.Time, day [34]bool, month [13]bool, weekdDay [8]bool) bool {
	dayOfMonth := nextDate.Day()
	monthOfYear := nextDate.Month()
	dayOfWeek := int(nextDate.Weekday())
	daysInCurrentMonth := daysInMonth(nextDate)

	if dayOfWeek == 0 {
		dayOfWeek = 7
	}

	switch t {
	case "w":
		return weekdDay[dayOfWeek]
	case "m":
		if day[predLastDayIndex] {
			if (dayOfMonth == daysInCurrentMonth-1) && month[monthOfYear] {
				return true
			}
		}

		if day[lastDayIndex] {
			if (dayOfMonth == daysInCurrentMonth) && month[monthOfYear] {
				return true
			}
		}

		return day[dayOfMonth] && month[monthOfYear]
	}

	return false
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if repeat == "" {
		return "", errors.New("repeat can't be empty")
	}

	date, err := time.Parse(timeLayout, dstart)

	if err != nil {
		return "", fmt.Errorf("failed to parse dstart: %w", err)
	}

	d, y := 0, 0

	var day [34]bool // 34 to handle -2, -1: day[33] is last day of month, day[32] is pred last day of month
	var month [13]bool
	var weekdDay [8]bool

	repeatParts := strings.Split(repeat, " ")

	switch repeatParts[0] {
	case "d":
		if len(repeatParts) != 2 {
			return "", errors.New("interval in days is not specified")
		}
		d, err = strconv.Atoi(repeatParts[1])
		if err != nil {
			return "", fmt.Errorf("invalid number of days: %w", err)
		}
		if d < 1 || d > 400 {
			return "", errors.New("invalid days interval: must be between 1 and 400")
		}
	case "y":
		y = 1
	case "w":
		if len(repeatParts) != 2 {
			return "", errors.New("interval in weeks is not specified")
		}

		splited := strings.Split(repeatParts[1], ",")

		if err := setRepeatWeekDays(&weekdDay, splited); err != nil {
			return "", fmt.Errorf("error setting repeat days of week: %w", err)
		}
	case "m":
		if len(repeatParts) < 2 {
			return "", errors.New("days of month are not specified")
		}

		splited := strings.Split(repeatParts[1], ",")

		if err := setRepeatDays(&day, splited); err != nil {
			return "", fmt.Errorf("error setting repeat days of month: %w", err)
		}

		if len(repeatParts) == 3 {
			splited = strings.Split(repeatParts[2], ",")
			if err := setRepeatMonths(&month, splited); err != nil {
				return "", fmt.Errorf("error setting repeat months: %w", err)
			}
		} else {
			for i := 1; i <= 12; i++ {
				month[i] = true
			}
		}
	default:
		return "", fmt.Errorf("invalid repeat format: %s", repeatParts[0])
	}

	repeatType := repeatParts[0]

	if repeatType == "d" || repeatType == "y" {
		for {
			date = date.AddDate(y, 0, d)
			if afterNow(date, now) {
				break
			}
		}

		return date.Format(timeLayout), nil
	}

	if afterNow(now, date) {
		date = now
	}

	for {
		date = date.AddDate(0, 0, 1)

		if checkDate(repeatType, date, day, month, weekdDay) {
			break
		}
	}

	return date.Format(timeLayout), nil
}

func nextDayHandler(w http.ResponseWriter, r *http.Request) {
	nowStr := r.FormValue("now")

	var nowTime time.Time
	var err error

	if nowStr == "" {
		nowTime = time.Now()
	} else {
		nowTime, err = time.Parse(timeLayout, nowStr)

		if err != nil {
			http.Error(w, "Invalid now time format", http.StatusBadRequest)
			return
		}
	}

	dstart := r.FormValue("date")
	if dstart == "" {
		http.Error(w, "Date is required", http.StatusBadRequest)
		return
	}

	repeat := r.FormValue("repeat")
	if repeat == "" {
		http.Error(w, "Repeat is required", http.StatusBadRequest)
		return
	}

	nextDate, err := NextDate(nowTime, dstart, repeat)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error in NextDate: %v", err), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(nextDate))
}
