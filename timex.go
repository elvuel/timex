package timex

import (
	"math"
	"time"
)

type timeFmt string

func (tf timeFmt) String() string {
	return string(tf)
}

const (
	// XYear format
	XYear timeFmt = "yy"
	// XMonth format
	XMonth timeFmt = "mm"
	// XDay format
	XDay timeFmt = "dd"
	// XHour format
	XHour timeFmt = "HH"
	// XMinute format
	XMinute timeFmt = "MM"
	// XSecond format
	XSecond timeFmt = "SS"

	// XWeek format
	XWeek timeFmt = "WK"
	// XSeason format
	XSeason timeFmt = "SMZ"
	// XSemiYear format
	XSemiYear timeFmt = "SMY"
)

// Weekday returns 1-7
func Weekday(t time.Time) int {
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return weekday
}

// Monday returns Monday of the week
func Monday(t time.Time) time.Time {
	return t.AddDate(0, 0, -Weekday(t)+1)
}

// Tuesday returns Tuesday of the week
func Tuesday(t time.Time) time.Time {
	return t.AddDate(0, 0, -Weekday(t)+2)
}

// Wednesday returns Wednesday of the week
func Wednesday(t time.Time) time.Time {
	return t.AddDate(0, 0, -Weekday(t)+3)
}

// Thursday returns Thursday of the week
func Thursday(t time.Time) time.Time {
	return t.AddDate(0, 0, -Weekday(t)+4)
}

// Friday returns Friday of the week
func Friday(t time.Time) time.Time {
	return t.AddDate(0, 0, -Weekday(t)+5)
}

// Saturday returns Saturday of the week
func Saturday(t time.Time) time.Time {
	return t.AddDate(0, 0, -Weekday(t)+6)
}

// Sunday returns Sunday of the week
func Sunday(t time.Time) time.Time {
	return t.AddDate(0, 0, -Weekday(t)+7)
}

// NextMonday returns next Monday of the week
func NextMonday(t time.Time) time.Time {
	return Monday(t).AddDate(0, 0, 7)
}

// NextTuesday returns next Tuesday of the week
func NextTuesday(t time.Time) time.Time {
	return Tuesday(t).AddDate(0, 0, 7)
}

// NextWednesday returns next Wednesday of the week
func NextWednesday(t time.Time) time.Time {
	return Wednesday(t).AddDate(0, 0, 7)
}

// NextThursday returns next Thursday of the week
func NextThursday(t time.Time) time.Time {
	return Thursday(t).AddDate(0, 0, 7)
}

// NextFriday returns next Friday of the week
func NextFriday(t time.Time) time.Time {
	return Friday(t).AddDate(0, 0, 7)
}

// NextSaturday returns next Saturday of the week
func NextSaturday(t time.Time) time.Time {
	return Saturday(t).AddDate(0, 0, 7)
}

// NextSunday returns next Sunday of the week
func NextSunday(t time.Time) time.Time {
	return Sunday(t).AddDate(0, 0, 7)
}

// LastMonday returns Last Monday of the week
func LastMonday(t time.Time) time.Time {
	return Monday(t).AddDate(0, 0, -7)
}

// LastTuesday returns Last Tuesday of the week
func LastTuesday(t time.Time) time.Time {
	return Tuesday(t).AddDate(0, 0, -7)
}

// LastWednesday returns Last Wednesday of the week
func LastWednesday(t time.Time) time.Time {
	return Wednesday(t).AddDate(0, 0, -7)
}

// LastThursday returns Last Thursday of the week
func LastThursday(t time.Time) time.Time {
	return Thursday(t).AddDate(0, 0, -7)
}

// LastFriday returns Last Friday of the week
func LastFriday(t time.Time) time.Time {
	return Friday(t).AddDate(0, 0, -7)
}

// LastSaturday returns Last Saturday of the week
func LastSaturday(t time.Time) time.Time {
	return Saturday(t).AddDate(0, 0, -7)
}

// LastSunday returns Last Sunday of the week
func LastSunday(t time.Time) time.Time {
	return Sunday(t).AddDate(0, 0, -7)
}

// BeginningOf returns beginning of the given descriptor format
func BeginningOf(t time.Time, dfmt string) time.Time {
	switch timeFmt(dfmt) {
	case XYear:
		y, _, _ := t.Date()
		return time.Date(y, time.January, 1, 0, 0, 0, 0, t.Location())
	case XMonth: // month
		y, m, _ := t.Date()
		return time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
	case XDay: // day
		y, m, d := t.Date()
		return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
	case XHour: // hour
		y, m, d := t.Date()
		return time.Date(y, m, d, t.Hour(), 0, 0, 0, t.Location())
	case XMinute: // minute
		return t.Truncate(time.Minute)

	case XWeek: // week
		return BeginningOf(t.AddDate(0, 0, -Weekday(t)+1), XDay.String())
	case XSeason: // season
		month := BeginningOf(t, XMonth.String())
		offset := (int(month.Month()) - 1) % 3
		return month.AddDate(0, -offset, 0)
	case XSemiYear: // semi
		month := BeginningOf(t, XMonth.String())
		offset := (int(month.Month()) - 1) % 6
		return month.AddDate(0, -offset, 0)
	}
	return t
}

// EndOf returns end of the given descriptor format
func EndOf(t time.Time, dfmt string) time.Time {
	switch timeFmt(dfmt) {
	case XYear:
		return BeginningOf(t, XYear.String()).AddDate(1, 0, 0).Add(-time.Nanosecond)
	case XMonth: // month
		return BeginningOf(t, XMonth.String()).AddDate(0, 1, 0).Add(-time.Nanosecond)
	case XDay: // day
		y, m, d := t.Date()
		return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
	case XHour: // hour
		return BeginningOf(t, XHour.String()).Add(time.Hour - time.Nanosecond)
	case XMinute: // minute
		return BeginningOf(t, XMinute.String()).Add(time.Minute - time.Nanosecond)

	case XWeek: // week
		return BeginningOf(t, XWeek.String()).AddDate(0, 0, 7).Add(-time.Nanosecond)
	case XSeason: // season
		return BeginningOf(t, XSeason.String()).AddDate(0, 3, 0).Add(-time.Nanosecond)
	case XSemiYear: // semi
		return BeginningOf(t, XSemiYear.String()).AddDate(0, 6, 0).Add(-time.Nanosecond)
	}
	return t
}

func absInterval(interval int) int {
	if interval <= 0 {
		interval = int(math.Abs(float64(interval)))
	}
	return interval
}

// XAt returns Last/Next x
func XAt(t time.Time, dfmt string, interval int) time.Time {
	switch timeFmt(dfmt) {
	case XYear:
		return t.AddDate(interval, 0, 0)
	case XMonth: // month
		return t.AddDate(0, interval, 0)
	case XDay: // day
		return t.AddDate(0, 0, interval)
	case XHour: // hour
		y, m, d := t.Date()
		return time.Date(y, m, d, t.Hour()+interval, t.Minute(), t.Second(), t.Nanosecond(), t.Location())
	case XMinute: // minute
		y, m, d := t.Date()
		return time.Date(y, m, d, t.Hour(), t.Minute()+interval, t.Second(), t.Nanosecond(), t.Location())
	case XSecond:
		y, m, d := t.Date()
		return time.Date(y, m, d, t.Hour(), t.Minute(), t.Second()+interval, t.Nanosecond(), t.Location())

	case XWeek: // x week[s] at same weekday
		return XAt(t, XDay.String(), interval*7)
	case XSeason: // x season[s] at 1st month and 1st week
		tx := BeginningOf(t, XSeason.String())
		tx = XAt(tx, XMonth.String(), interval*3)
		y, m, d := tx.Date()
		return time.Date(y, m, d+Weekday(t)-1, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
	case XSemiYear: // x semi-year[s] at 1st month and 1st week
		tx := BeginningOf(t, XSemiYear.String())
		tx = XAt(tx, XMonth.String(), interval*6)
		y, m, d := tx.Date()
		return time.Date(y, m, d+Weekday(t)-1, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
	}

	return t
}

// LastXAt returns last time with the given descriptor format
func LastXAt(t time.Time, dfmt string, interval int) time.Time {
	interval = absInterval(interval) * -1
	return XAt(t, dfmt, interval)
}

// NextXAt returns last time with the given descriptor format
func NextXAt(t time.Time, dfmt string, interval int) time.Time {
	return XAt(t, dfmt, absInterval(interval))
}
