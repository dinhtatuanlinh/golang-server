package utils

import (
	"math"
	"time"
)

func Now(hourOffset int, format string) string {
	return time.Now().UTC().Add(time.Hour * time.Duration(hourOffset)).Format(format)
}

func GetTime(hourOffset int) time.Time {
	return time.Now().UTC().Add(time.Hour * time.Duration(hourOffset))
}

func GetSpecificTime(hourOffset int, minuteOffset int, secondOffset int) time.Time {
	return time.Now().UTC().Add(time.Hour * time.Duration(hourOffset)).Add(time.Minute * time.Duration(minuteOffset)).Add(time.Second * time.Duration(secondOffset))
}

func GetTimeDifference(now time.Time, then time.Time) time.Duration {
	return now.Sub(then)
}

func GetHourDifference(now time.Time, then time.Time) int {
	return int(math.Round(now.Sub(then).Hours()))
}

func GetMinuteDifference(now time.Time, then time.Time) int {
	return int(math.Round(now.Sub(then).Minutes()))
}

func GetSecondDifference(now time.Time, then time.Time) int {
	return int(math.Round(now.Sub(then).Seconds()))
}

// GetTimeFromTimestamp is a func to get time in iso format
func GetTimeFromTimestamp(timestamp int) time.Time {
	return time.Unix(int64(timestamp), 0)
}

func GetDifferenceTime(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
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

	return
}

func NextMonday(currentDate time.Time) time.Time {
	var nextDayToMonday int
	switch currentDate.Weekday() {
	case time.Sunday:
		nextDayToMonday = 1
	case time.Saturday:
		nextDayToMonday = 2
	case time.Friday:
		nextDayToMonday = 3
	case time.Thursday:
		nextDayToMonday = 4
	case time.Wednesday:
		nextDayToMonday = 5
	case time.Tuesday:
		nextDayToMonday = 6
	case time.Monday:
		nextDayToMonday = 7
	}
	return currentDate.AddDate(0, 0, nextDayToMonday)
}

func NextDay(currentDate time.Time) time.Time {
	return currentDate.AddDate(0, 0, 1)
}
