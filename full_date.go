package timecalc

import (
	"math"
	"time"
)

type FullDate struct {
	*time.Time
}

type DateDiff struct {
	Years  int
	Months int
	Days   int
}

type Duration struct {
	Years      int
	Months     int
	Days       int
	IsFinished bool
	Error      error
}

// Age returns the number of years, months, and days passed at the given
// reference date since this FullDate.
func (fd FullDate) Age(referenceDate time.Time) DateDiff {
	referenceDate = stripTimeComponent(referenceDate)
	return GetDateDiff(fd, FullDate{&referenceDate})
}

// AgeYears returns the number of full years passed for the given reference date
// since this FullDate.
func (fd FullDate) AgeYears(referenceDate time.Time) int {
	return fd.Age(referenceDate).Years
}

func (fd FullDate) InverseAge(referenceDate time.Time) DateDiff {
	referenceDate = stripTimeComponent(referenceDate)
	return GetDateDiff(FullDate{&referenceDate}, fd)
}

func (fd FullDate) InverseAgeYears(referenceDate time.Time) int {
	return fd.InverseAge(referenceDate).Years
}

func (fd FullDate) WeekdayName() string {
	switch fd.Weekday() {
	case time.Sunday:
		return "Sonntag"
	case time.Monday:
		return "Montag"
	case time.Tuesday:
		return "Diestag"
	case time.Wednesday:
		return "Mittwoch"
	case time.Thursday:
		return "Donnerstag"
	case time.Friday:
		return "Freitag"
	case time.Saturday:
		return "Samstag"
	}
	return ""
}

func (fd FullDate) NextInDays(referenceDate time.Time) int {
	referenceDate = stripTimeComponent(referenceDate)
	hoursToGo := fd.NextAnniversary(referenceDate).Sub(referenceDate).Hours()
	return int(math.Floor(hoursToGo / 24))
}

func (fd FullDate) NextAnniversary(referenceDate time.Time) time.Time {
	referenceDate = stripTimeComponent(referenceDate)
	nextDate := time.Date(fd.Year(), fd.Month(), fd.Day(), fd.Hour(), fd.Minute(), fd.Second(), fd.Nanosecond(), fd.Location())
	nextDate = nextDate.AddDate(referenceDate.Year()-fd.Year(), 0, 0)
	if referenceDate.Month() > fd.Month() || (referenceDate.Month() == fd.Month() && referenceDate.Day() > fd.Day()) {
		nextDate = nextDate.AddDate(1, 0, 0)
	}
	return nextDate
}

func (fd FullDate) NextAnniversaryFullDate(referenceDate time.Time) FullDate {
	nextDate := fd.NextAnniversary(referenceDate)
	return FullDate{&nextDate}
}

func GetDateDiff(a, b FullDate) DateDiff {
	years := b.Year() - a.Year()
	if b.Month() < a.Month() ||
		(b.Month() == a.Month() && b.Day() < a.Day()) {
		years--
	}

	months := int(b.Month()) - int(a.Month())
	if b.Day() < a.Day() {
		months--
	}
	if months < 0 {
		months += 12
	}

	var days int
	if b.Day() >= a.Day() {
		days = b.Day() - a.Day()
	} else {
		days = (getMonthDays(a.Year(), int(a.Month())) - a.Day()) + b.Day()
	}

	return DateDiff{
		Years:  years,
		Months: months,
		Days:   days,
	}
}

func stripTimeComponent(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
