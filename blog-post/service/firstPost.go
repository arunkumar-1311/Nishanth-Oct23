package service

import (
	"fmt"
	"time"
)

// Helps to find the time difference between 2 time period
func TimeDifference(date, now time.Time, timeDuration *string) error {

	if date.Location() != now.Location() {
		now = now.In(date.Location())
	}
	if date.After(now) {
		date, now = now, date
	}
	y1, M1, d1 := date.Date()
	y2, M2, d2 := now.Date()

	year := int(y2 - y1)
	month := int(M2 - M1)
	day := int(d2 - d1)

	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}
	*timeDuration = fmt.Sprintf("%vy %vm %vd", year, month, day)
	return nil
}
