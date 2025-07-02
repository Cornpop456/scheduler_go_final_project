package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func afterNow(date, now time.Time) bool {
	dy, dm, dd := date.Date()
	ny, nm, nd := now.Date()

	if dy != ny {
		return dy > ny
	}

	if dm != nm {
		return dm > nm
	}

	return dd > nd
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
	default:
		return "", fmt.Errorf("invalid repeat format: %s", repeatParts[0])
	}

	for {
		date = date.AddDate(y, 0, d)
		if afterNow(date, now) {
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
