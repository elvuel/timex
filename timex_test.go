package timex

import (
	"math"
	"testing"
	"time"
)

type ttime func(time.Time) time.Time

func TestAll_Weekdays(t *testing.T) {
	now := time.Now()
	var tx time.Time

	weekdays := make([][]time.Time, 8)
	for i := 1; i <= 7; i++ {
		weekdays[i-1] = make([]time.Time, 0)
	}
	for i := -366; i <= 366; i++ {
		tx = now.AddDate(0, 0, i)
		wd := tx.Weekday()
		twd := Weekday(tx)
		if wd == 0 {
			if twd != 7 {
				t.Error("should be sunday")
			}
		} else {
			if int(wd) != twd {
				t.Errorf("should got `%s`, but got `%s`\n", wd.String(), time.Weekday(twd).String())
			}
		}
		weekdays[twd] = append(weekdays[twd], tx)
	}

	var tf ttime
	for i := 1; i <= 6; i++ {
		switch i {
		case 1:
			tf = Monday
		case 2:
			tf = Tuesday
		case 3:
			tf = Wednesday
		case 4:
			tf = Thursday
		case 5:
			tf = Friday
		case 6:
			tf = Saturday
		}
		for _, tm := range weekdays[i] {
			if v := int(tf(tm).Weekday()); v != i {
				t.Errorf("should be %d, index %d\n", v, i)
			}
		}
		weekdays[i] = weekdays[i][:0]
	}
	tf = Sunday
	for _, tm := range weekdays[7] {
		if v := int(tf(tm).Weekday()); v != 0 {
			t.Errorf("should be %d, but got %d\n", 0, v)
		}
	}
}

func TestAll_NLWeekdays(t *testing.T) {
	now, _ := time.Parse("2006-01-02", "2019-10-24")

	fmtf := func(tx time.Time) string {
		return tx.Format("2006-01-02")
	}

	funcs := make([][]ttime, 8)
	funcs[0] = []ttime{
		Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday,
		LastMonday, LastTuesday, LastWednesday, LastThursday, LastFriday, LastSaturday, LastSunday,
		NextMonday, NextTuesday, NextWednesday, NextThursday, NextFriday, NextSaturday, NextSunday,
	}
	funcs[1] = []ttime{Monday, LastMonday, NextMonday}
	funcs[2] = []ttime{Tuesday, LastTuesday, NextTuesday}
	funcs[3] = []ttime{Wednesday, LastWednesday, NextWednesday}
	funcs[4] = []ttime{Thursday, LastThursday, NextThursday}
	funcs[5] = []ttime{Friday, LastFriday, NextFriday}
	funcs[6] = []ttime{Saturday, LastSaturday, NextSaturday}
	funcs[7] = []ttime{Sunday, LastSunday, NextSunday}

	seeds := make([][]string, 8)
	seeds[0] = []string{
		"10-21", "10-22", "10-23", "10-24", "10-25", "10-26", "10-27",
		"10-14", "10-15", "10-16", "10-17", "10-18", "10-19", "10-20",
		"10-28", "10-29", "10-30", "10-31", "11-01", "11-02", "11-03",
	}
	seeds[1] = []string{"10-21", "10-14", "10-28"}
	seeds[2] = []string{"10-22", "10-15", "10-29"}
	seeds[3] = []string{"10-23", "10-16", "10-30"}
	seeds[4] = []string{"10-24", "10-17", "10-31"}
	seeds[5] = []string{"10-25", "10-18", "11-01"}
	seeds[6] = []string{"10-26", "10-19", "11-02"}
	seeds[7] = []string{"10-27", "10-20", "11-03"}
	for i := -3; i <= 3; i++ {
		tx := now.AddDate(0, 0, i)
		wd := Weekday(tx)
		for j := 0; j < 3; j++ {
			if dfmt, rfmt := fmtf(funcs[wd][j](tx)), ("2019-" + seeds[wd][j]); dfmt != rfmt {
				t.Errorf("want %s but got %s\n", rfmt, dfmt)
			}
		}
	}

	for i, fk := range funcs[0] {
		if dfmt, rfmt := fmtf(fk(now)), "2019-"+seeds[0][i]; dfmt != rfmt {
			t.Errorf("want %s but got %s\n", rfmt, dfmt)
		}
	}
}

func TestXFmt(t *testing.T) {
	fmts := []timeFmt{XYear, XMonth, XDay, XHour, XMinute, XSecond, XWeek, XSeason, XSemiYear}
	for i, dfmt := range fmts {
		if dfmt.String() != string(fmts[i]) {
			t.Errorf("should be all right fmts")
		}
	}
}

func TestBeginingOf(t *testing.T) {
	now, _ := time.Parse("2006-01-02 15:04:05", "2019-10-24 15:04:05")

	fmtf := func(tx time.Time) string {
		return tx.Format(time.RFC3339)
	}

	fmts := []timeFmt{XYear, XMonth, XDay, XHour, XMinute, XWeek, XSeason, XSemiYear}

	seeds := map[string]string{
		XHour.String():     "2019-10-24T15:00:00Z",
		XMinute.String():   "2019-10-24T15:04:00Z",
		XSemiYear.String(): "2019-07-01T00:00:00Z",
		XSeason.String():   "2019-10-01T00:00:00Z",
		XWeek.String():     "2019-10-21T00:00:00Z",
		XDay.String():      "2019-10-24T00:00:00Z",
		XMonth.String():    "2019-10-01T00:00:00Z",
		XYear.String():     "2019-01-01T00:00:00Z",
	}
	for _, dfmt := range fmts {
		if seeds[dfmt.String()] != fmtf(BeginningOf(now, dfmt.String())) {
			t.Error("should be the same")
		}
	}

	if fmtf(BeginningOf(now, "foobar")) != fmtf(now) {
		t.Error("should be return the same time with unkown format")
	}
}

func TestEndOf(t *testing.T) {
	now, _ := time.Parse("2006-01-02 15:04:05", "2019-10-24 15:04:05")

	fmtf := func(tx time.Time) string {
		return tx.Format(time.RFC3339)
	}

	fmts := []timeFmt{XYear, XMonth, XDay, XHour, XMinute, XWeek, XSeason, XSemiYear}

	seeds := map[string]string{
		XHour.String():     "2019-10-24T15:59:59Z",
		XMinute.String():   "2019-10-24T15:04:59Z",
		XSemiYear.String(): "2019-12-31T23:59:59Z",
		XSeason.String():   "2019-12-31T23:59:59Z",
		XWeek.String():     "2019-10-27T23:59:59Z",
		XDay.String():      "2019-10-24T23:59:59Z",
		XMonth.String():    "2019-10-31T23:59:59Z",
		XYear.String():     "2019-12-31T23:59:59Z",
	}

	for _, dfmt := range fmts {
		if seeds[dfmt.String()] != fmtf(EndOf(now, dfmt.String())) {
			t.Error("should be the same")
		}
	}

	if fmtf(EndOf(now, "foobar")) != fmtf(now) {
		t.Error("should be return the same time with unkown format")
	}
}

func TestAbsInterval(t *testing.T) {
	for i := -1; i <= 1; i++ {
		j := int(math.Abs(float64(i)))
		if absInterval(i) != j {
			t.Error("should be return the right abs value")
		}
	}

}

func TestXAt(t *testing.T) {
	now, _ := time.Parse("2006-01-02 15:04:05", "2019-10-24 15:04:05")

	fmtf := func(tx time.Time) string {
		return tx.Format(time.RFC3339)
	}

	// Next
	nextSeeds := map[string]string{
		XHour.String():     "2019-10-24T16:04:05Z",
		XMinute.String():   "2019-10-24T15:05:05Z",
		XSemiYear.String(): "2020-01-04T15:04:05Z",
		XSeason.String():   "2020-01-04T15:04:05Z",
		XSecond.String():   "2019-10-24T15:04:06Z",
		XWeek.String():     "2019-10-31T15:04:05Z",
		XDay.String():      "2019-10-25T15:04:05Z",
		XMonth.String():    "2019-11-24T15:04:05Z",
		XYear.String():     "2020-10-24T15:04:05Z",
	}

	// Last
	lastSeeds := map[string]string{
		XHour.String():     "2019-10-24T14:04:05Z",
		XMinute.String():   "2019-10-24T15:03:05Z",
		XSemiYear.String(): "2019-01-04T15:04:05Z",
		XSeason.String():   "2019-07-04T15:04:05Z",
		XSecond.String():   "2019-10-24T15:04:04Z",
		XWeek.String():     "2019-10-17T15:04:05Z",
		XDay.String():      "2019-10-23T15:04:05Z",
		XMonth.String():    "2019-09-24T15:04:05Z",
		XYear.String():     "2018-10-24T15:04:05Z",
	}

	fmts := []timeFmt{XYear, XMonth, XDay, XHour, XMinute, XSecond, XWeek, XSeason, XSemiYear}
	for _, dfmt := range fmts {
		if nextSeeds[dfmt.String()] != fmtf(XAt(now, dfmt.String(), 1)) || nextSeeds[dfmt.String()] != fmtf(NextXAt(now, dfmt.String(), 1)) {
			t.Error("should be the same nexts")
		}
		if lastSeeds[dfmt.String()] != fmtf(XAt(now, dfmt.String(), -1)) || lastSeeds[dfmt.String()] != fmtf(LastXAt(now, dfmt.String(), 1)) {
			t.Error("should be the same lasts")
		}
	}
	if fmtf(XAt(now, "foobar", 1)) != fmtf(now) {
		t.Error("should be return the same time with unkown format")
	}

	if fmtf(LastXAt(now, "foobar", 1)) != fmtf(now) {
		t.Error("should be return the same time with unkown format")
	}

	if fmtf(NextXAt(now, "foobar", 1)) != fmtf(now) {
		t.Error("should be return the same time with unkown format")
	}
}
