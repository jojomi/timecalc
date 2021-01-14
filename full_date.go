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
	return CalcDateDiff(fd, FullDate{&referenceDate})
}

// AgeYears returns the number of full years passed for the given reference date
// since this FullDate.
func (fd FullDate) AgeYears(referenceDate time.Time) int {
	return fd.Age(referenceDate).Years
}

func (fd FullDate) InverseAge(referenceDate time.Time) DateDiff {
	referenceDate = stripTimeComponent(referenceDate)
	return CalcDateDiff(FullDate{&referenceDate}, fd)
}

func (fd FullDate) InverseAgeYears(referenceDate time.Time) int {
	return fd.InverseAge(referenceDate).Years
}

type AstrologicalSign struct {
	Name          string
	LocalizedName string
	Symbol        string
}

var astrologicalSigns = map[string]AstrologicalSign{
	"Aries": AstrologicalSign{
		Name:          "Aries",
		LocalizedName: "Widder",
		Symbol:        "♈",
	},
	"Taurus": AstrologicalSign{
		Name:          "Taurus",
		LocalizedName: "Stier",
		Symbol:        "♉",
	},
	"Gemini": AstrologicalSign{
		Name:          "Gemini",
		LocalizedName: "Zwillinge",
		Symbol:        "♊",
	},
	"Cancer": AstrologicalSign{
		Name:          "Cancer",
		LocalizedName: "Krebs",
		Symbol:        "♋",
	},
	"Leo": AstrologicalSign{
		Name:          "Leo",
		LocalizedName: "Löwe",
		Symbol:        "♌",
	},
	"Virgo": AstrologicalSign{
		Name:          "Virgo",
		LocalizedName: "Jungfrau",
		Symbol:        "♍",
	},
	"Libra": AstrologicalSign{
		Name:          "Libra",
		LocalizedName: "Waage",
		Symbol:        "♎",
	},
	"Scorpio": AstrologicalSign{
		Name:          "Scorpio",
		LocalizedName: "Skorpion",
		Symbol:        "♏",
	},
	"Sagittarius": AstrologicalSign{
		Name:          "Sagittarius",
		LocalizedName: "Schütze",
		Symbol:        "♐",
	},
	"Capricornus": AstrologicalSign{
		Name:          "Capricornus",
		LocalizedName: "Steinbock",
		Symbol:        "♑",
	},
	"Aquarius": AstrologicalSign{
		Name:          "Aquarius",
		LocalizedName: "Wassermann",
		Symbol:        "♒",
	},
	"Pisces": AstrologicalSign{
		Name:          "Pisces",
		LocalizedName: "Fische",
		Symbol:        "♓",
	},
}

func (fd FullDate) AstrologicalSign() AstrologicalSign {
	switch fd.Month() {
	case time.January:
		if fd.Day() <= 20 {
			return makeAstrologicalSign("Capricornus")
		}
		return makeAstrologicalSign("Aquarius")
	case time.February:
		if fd.Day() <= 19 {
			return makeAstrologicalSign("Aquarius")
		}
		return makeAstrologicalSign("Pisces")
	case time.March:
		if fd.Day() <= 20 {
			return makeAstrologicalSign("Pisces")
		}
		return makeAstrologicalSign("Aries")
	case time.April:
		if fd.Day() <= 20 {
			return makeAstrologicalSign("Aries")
		}
		return makeAstrologicalSign("Taurus")
	case time.May:
		if fd.Day() <= 20 {
			return makeAstrologicalSign("Taurus")
		}
		return makeAstrologicalSign("Gemini")
	case time.June:
		if fd.Day() <= 20 {
			return makeAstrologicalSign("Gemini")
		}
		return makeAstrologicalSign("Cancer")
	case time.July:
		if fd.Day() <= 22 {
			return makeAstrologicalSign("Cancer")
		}
		return makeAstrologicalSign("Leo")
	case time.August:
		if fd.Day() <= 23 {
			return makeAstrologicalSign("Leo")
		}
		return makeAstrologicalSign("Virgo")
	case time.September:
		if fd.Day() <= 23 {
			return makeAstrologicalSign("Virgo")
		}
		return makeAstrologicalSign("Libra")
	case time.October:
		if fd.Day() <= 23 {
			return makeAstrologicalSign("Libra")
		}
		return makeAstrologicalSign("Scorpio")
	case time.November:
		if fd.Day() <= 22 {
			return makeAstrologicalSign("Scorpio")
		}
		return makeAstrologicalSign("Sagittarius")
	case time.December:
		if fd.Day() <= 21 {
			return makeAstrologicalSign("Sagittarius")
		}
		return makeAstrologicalSign("Capricornus")
	}
	return AstrologicalSign{}
}

func makeAstrologicalSign(name string) AstrologicalSign {
	if astrologicalSign, ok := astrologicalSigns[name]; ok {
		return astrologicalSign
	}
	return AstrologicalSign{}
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

func CalcDateDiff(a, b FullDate) DateDiff {
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
