package timecalc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSparseDateFrom(t *testing.T) {
	// parse a full date
	sd1 := SparseDateFrom("1924-12-13")
	assert.Equal(t, "1924-12-13", sd1.GetDate().Format("2006-01-02"))

	// parse incomplete (sparse) date
	sd2 := SparseDateFrom("1924-12-??")
	assert.True(t, sd2.GetDate().IsZero())
}

func TestSparseDateFromLocalized(t *testing.T) {
	// parse full localized date (custom format)
	sd1 := SparseDateFromLocalized("1924.12.14", "2006.01.02")
	assert.Equal(t, "1924-12-14", sd1.GetDate().Format("2006-01-02"))

	// parse incomplete localized date (custom format)
	sd2 := SparseDateFromLocalized("1924.12.??", "2006.01.02")
	assert.True(t, sd2.GetDate().IsZero())
}

func TestComplete(t *testing.T) {
	sd1 := SparseDateFrom("1924-12-15")
	assert.True(t, sd1.IsComplete(), "SparseDate is complete")

	sd2 := SparseDateFrom("1924-12-??")
	assert.False(t, sd2.IsComplete(), "SparseDate is not complete")
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		Input    string
		Expected bool
	}{
		{
			Input:    "????-??-??",
			Expected: true,
		},
		{
			Input:    "1985-06-??",
			Expected: false,
		},
		{
			Input:    "1983-01-02",
			Expected: false,
		},
	}
	var sd *SparseDate
	for _, test := range tests {
		sd = SparseDateFrom(test.Input)
		assert.Equal(t, test.Expected, sd.IsEmpty(), fmt.Sprintf("IsEmpty for %s should be %t", test.Input, test.Expected))
	}
}

type InOutTest struct {
	in, out string
}

func TestGetMinDate(t *testing.T) {
	var sd *SparseDate
	tests := []InOutTest{
		{
			in:  "1975-06-??",
			out: "1975-06-01",
		},
		{
			in:  "1966-??-04",
			out: "1966-01-04",
		},
		{
			in:  "1982-??-??",
			out: "1982-01-01",
		},
	}
	for _, test := range tests {
		sd = SparseDateFrom(test.in)
		assert.Equal(t, test.out, sd.GetMinDate().Format("2006-01-02"), fmt.Sprintf("MinDate for %s should be %s", test.in, test.out))
	}
}

func TestGetMaxDate(t *testing.T) {
	var sd *SparseDate
	tests := []InOutTest{
		{
			in:  "1975-06-??",
			out: "1975-06-30",
		},
		{
			in:  "2?0?-02-??",
			out: "2909-02-28",
		},
		{
			in:  "2000-02-??",
			out: "2000-02-29",
		},
		{
			in:  "2100-02-??",
			out: "2100-02-28",
		},
		{
			in:  "2004-02-??",
			out: "2004-02-29",
		},
		{
			in:  "1966-??-04",
			out: "1966-12-04",
		},
		{
			in:  "1982-??-??",
			out: "1982-12-31",
		},
	}
	for _, test := range tests {
		sd = SparseDateFrom(test.in)
		assert.Equal(t, test.out, sd.GetMaxDate().Format("2006-01-02"), fmt.Sprintf("MaxDate for %s should be %s", test.in, test.out))
	}
}

func TestFormat(t *testing.T) {
	tests := []struct {
		input    string
		format   string
		expected string
	}{
		{"2010-05-06", "02.01.2006", "06.05.2010"},
		{"2010-05-06", "2006", "2010"},
		{"20??-05-??", "02.01.2006", "??.05.20??"},
		{"2011-??-??", "02.01.2006", "2011"},
		{"????-??-??", "02.01.2006", ""},
	}

	for _, test := range tests {
		sd1 := SparseDateFrom(test.input)
		assert.Equal(t, test.expected, sd1.Format(test.format))
	}
}
