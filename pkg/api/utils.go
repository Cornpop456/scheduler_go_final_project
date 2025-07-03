package api

import "time"

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

func daysInMonth(t time.Time) int {
	firstOfMonth := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	firstOfNextMonth := firstOfMonth.AddDate(0, 1, 0)
	lastOfMonth := firstOfNextMonth.AddDate(0, 0, -1)
	return lastOfMonth.Day()
}
