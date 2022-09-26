package util

import (
	"github.com/robfig/cron"
	"regexp"
	"time"
)

const (
	// Set the top bit if a star was included in the expression.
	starBit = 1 << 63
)

var fieldFinder = regexp.MustCompile(`\S+`)

// function CronCompatibility is used to compatible with both 5 field and 6 field cron
func CronCompatibility(cron string) string {
	fields := fieldFinder.FindAllString(cron, -1)
	if len(fields) == 5 {
		cron = "0 " + cron
	}
	return cron
}

func CronPre(t time.Time, s cron.SpecSchedule) time.Time {

	// Start at the earliest possible time (the upcoming second).
	t = t.Add(1*time.Second - time.Duration(t.Nanosecond())*time.Nanosecond)

	// This flag indicates whether a field has been incremented.
	added := false

	// If no time is found within five years, return zero.

	dayMatchesFunc := func(s cron.SpecSchedule, t time.Time) bool {
		var (
			domMatch bool = 1<<uint(t.Day())&s.Dom > 0
			dowMatch bool = 1<<uint(t.Weekday())&s.Dow > 0
		)
		if s.Dom&starBit > 0 || s.Dow&starBit > 0 {
			return domMatch && dowMatch
		}
		return domMatch || dowMatch
	}
	yearLimit := t.Year() - 5

WRAP:
	if t.Year() < yearLimit {
		return time.Time{}
	}

	for 1<<uint(t.Second())&s.Second == 0 {
		if !added {
			added = true
			t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, t.Location())
		}
		t = t.Add(-1 * time.Second)

		if t.Second() == 0 {
			goto WRAP
		}
	}

	for 1<<uint(t.Minute())&s.Minute == 0 {
		if !added {
			added = true
			t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, t.Location())
		}
		t = t.Add(-1 * time.Minute)

		if t.Minute() == 0 {
			goto WRAP
		}
	}

	for 1<<uint(t.Hour())&s.Hour == 0 {
		if !added {
			added = true
			t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
		}
		t = t.Add(-1 * time.Hour)

		if t.Hour() == 0 {
			goto WRAP
		}
	}

	// Now get a day in that month.
	for !dayMatchesFunc(s, t) {
		if !added {
			added = true
			t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
		}
		t = t.AddDate(0, 0, -1)

		if t.Day() == 1 {
			goto WRAP
		}
	}

	// Find the first applicable month.
	// If it's this month, then do nothing.
	for 1<<uint(t.Month())&s.Month == 0 {
		// If we have to add a month, reset the other parts to 0.
		if !added {
			added = true
			// Otherwise, set the date at the beginning (since the current time is irrelevant).
			t = time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
		}
		t = t.AddDate(0, -1, 0)

		// Wrapped around.
		if t.Month() == time.January {
			goto WRAP
		}
	}

	return t
}
