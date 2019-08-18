package timecalc

import (
	"strconv"
	"strings"
	"time"
)

var monthDaysMap = map[int]int{
	1:  31,
	2:  28,
	3:  31,
	4:  30,
	5:  31,
	6:  30,
	7:  31,
	8:  31,
	9:  30,
	10: 31,
	11: 30,
	12: 31,
}

func getMonthDays(year, month int) int {
	if month == 2 && isLeapYear(year) {
		return 29
	}
	return monthDaysMap[month]
}

type SparseDate struct {
	// Any of the y's is allowed to be replaced by ? as well as mm and/or dd.
	// Any combination is valid: 196?-03-?, 20??-05-08, 20??-?-?, 20??-?-?, 1985-12-31, ????-?-?
	year, month, day string
}

func SparseDateFrom(date string) *SparseDate {
	return SparseDateFromLocalized(date, "2006-01-02")
}

func SparseDateFromLocalized(input, dateFormat string) *SparseDate {
	s := &SparseDate{}
	s.Parse(input, dateFormat)
	return s
}

func (sd *SparseDate) IsComplete() bool {
	return strings.Index(sd.year, "?") == -1 &&
		strings.Index(sd.month, "?") == -1 &&
		strings.Index(sd.day, "?") == -1
}

func (sd *SparseDate) Year() string {
	return sd.year
}

func (sd *SparseDate) YearInt() int {
	if !sd.IsCompleteYear() {
		return 0
	}
	year, err := strconv.Atoi(sd.Year())
	if err != nil {
		return 0
	}
	return year
}

func (sd *SparseDate) IsCompleteYear() bool {
	return strings.Index(sd.year, "?") == -1
}

func (sd *SparseDate) IsCompleteDayAndMonth() bool {
	return strings.Index(sd.day, "?") == -1 && strings.Index(sd.month, "?") == -1
}

func (sd *SparseDate) GetMinDate() time.Time {
	// check year
	outputYear := strings.Replace(sd.year, "?", "0", -1)
	// check month
	outputMonth := sd.month
	if outputMonth == "?" || outputMonth == "??" {
		outputMonth = "01"
	}
	// check day
	outputDay := sd.day
	if outputDay == "?" || outputDay == "??" {
		outputDay = "01"
	}
	time, _ := time.Parse("20060102", outputYear+outputMonth+outputDay)
	return time
}

func (sd *SparseDate) GetMin() FullDate {
	min := sd.GetMinDate()
	return FullDate{&min}
}

func (sd *SparseDate) GetMaxDate() time.Time {
	// check year
	outputYear := strings.Replace(sd.year, "?", "9", -1)
	// check month
	outputMonth := sd.month
	if outputMonth == "?" || outputMonth == "??" {
		outputMonth = "12"
	}
	// check day
	outputDay := sd.day
	if outputDay == "?" || outputDay == "??" {
		month, err := strconv.Atoi(outputMonth)
		if err == nil {
			outputYearInt, errInt := strconv.Atoi(outputYear)
			if errInt == nil {
				outputDayInt := getMonthDays(outputYearInt, month)
				outputDay = strconv.Itoa(outputDayInt)
			}
		}
	}
	time, _ := time.Parse("20060102", outputYear+outputMonth+outputDay)
	return time
}

func (sd *SparseDate) GetMax() FullDate {
	max := sd.GetMaxDate()
	return FullDate{&max}
}

func (sd *SparseDate) GetDate() time.Time {
	if !sd.IsComplete() {
		return time.Time{}
	}
	date, err := time.Parse("20060102", sd.year+sd.month+sd.day)
	if err != nil {
		return time.Time{}
	}
	return date
}

func (sd *SparseDate) Date() FullDate {
	date := sd.GetDate()
	return FullDate{&date}
}

func (sd *SparseDate) GetDateFormatted(format string) string {
	date := sd.GetDate()
	if date.IsZero() {
		return ""
	}
	return date.Format(format)
}

func (sd *SparseDate) Parse(input, dateFormat string) {
	yearIndex := strings.Index(dateFormat, "2006")
	if len(input) >= yearIndex+4 {
		sd.year = input[yearIndex : yearIndex+4]
	}
	monthIndex := strings.Index(dateFormat, "01")
	if len(input) >= monthIndex+2 {
		sd.month = input[monthIndex : monthIndex+2]
	}
	dayIndex := strings.Index(dateFormat, "02")
	if len(input) >= dayIndex+2 {
		sd.day = input[dayIndex : dayIndex+2]
	}
}

func (sd *SparseDate) IsEmpty() bool {
	return sd.year == "" && sd.month == "" && sd.day == ""
}

func (sd *SparseDate) String() string {
	if sd.IsEmpty() {
		return ""
	}
	return sd.year + "-" + sd.month + "-" + sd.day
}

func (sd *SparseDate) FormatShort(formatString string) string {
	// only year set?
	if sd.IsCompleteYear() && sd.day == "??" && sd.month == "??" {
		return sd.year
	}
	return sd.Format(formatString)
}

func (sd *SparseDate) Format(formatString string) string {
	if sd.IsEmpty() {
		return ""
	}
	var result string
	// special case: only year given
	if sd.month == "??" && sd.day == "??" {
		result = "2006"
	} else {
		result = formatString
	}
	replacer := strings.NewReplacer("2006", sd.year, "01", sd.month, "02", sd.day)
	result = replacer.Replace(result)
	return result
}

func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func (sd SparseDate) MarshalJSON() ([]byte, error) {
	return []byte(`"` + sd.Format("2006-01-02") + `"`), nil
}
