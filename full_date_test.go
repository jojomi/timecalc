package timecalc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAge(t *testing.T) {
	tests := []struct {
		dateA    time.Time
		dateB    time.Time
		expected DateDiff
	}{
		{
			time.Date(2018, 10, 31, 15, 45, 58, 0, time.UTC),
			time.Date(2010, 05, 06, 0, 0, 0, 0, time.UTC),
			DateDiff{Years: 8, Months: 5, Days: 25},
		},
		// special case: year overflow (monath!)
		{
			time.Date(2018, 10, 15, 15, 45, 58, 0, time.UTC),
			time.Date(2017, 12, 15, 0, 0, 0, 0, time.UTC),
			DateDiff{Years: 0, Months: 10, Days: 0},
		},
		// month overflow (days!)
		{
			time.Date(2018, 10, 05, 15, 45, 58, 0, time.UTC),
			time.Date(2013, 05, 15, 0, 0, 0, 0, time.UTC),
			DateDiff{Years: 5, Months: 4, Days: 21},
		},
		// special case: same day
		{
			time.Date(2018, 10, 05, 15, 45, 58, 0, time.UTC),
			time.Date(2018, 10, 05, 0, 0, 0, 0, time.UTC),
			DateDiff{Years: 0, Months: 0, Days: 0},
		},
	}

	for _, test := range tests {
		age := FullDate{&test.dateB}.Age(test.dateA)
		assert.Equal(t, test.expected.Years, age.Years, "Age Years check")
		assert.Equal(t, test.expected.Months, age.Months, "Age Months check")
		assert.Equal(t, test.expected.Days, age.Days, "Age Days check")
	}
}

func TestAgeYears(t *testing.T) {
	referenceDate := time.Date(2015, 10, 31, 15, 45, 58, 0, time.UTC)
	full := time.Date(2010, 05, 06, 0, 0, 0, 0, time.UTC)
	sd1 := FullDate{&full}
	age := sd1.AgeYears(referenceDate)
	assert.Equal(t, 5, age)
}

func TestAgeCornerCase(t *testing.T) {
	birth := time.Date(2012, 04, 01, 0, 0, 0, 0, time.UTC)
	bdate := FullDate{&birth}

	now := time.Date(2015, 04, 01, 0, 0, 1, 0, time.UTC)
	age := bdate.AgeYears(now)
	assert.Equal(t, 3, age)

	now2 := time.Date(2015, 03, 31, 23, 59, 59, 0, time.UTC)
	age2 := bdate.AgeYears(now2)
	assert.Equal(t, 2, age2)
}

func TestNextAnniversary(t *testing.T) {
	referenceDate := time.Date(2015, 10, 31, 15, 45, 58, 0, time.UTC)

	full := time.Date(2010, 05, 06, 0, 0, 0, 0, time.UTC)
	sd1 := FullDate{&full}
	next := sd1.NextAnniversary(referenceDate)
	assert.Equal(t, "2016-05-06", next.Format("2006-01-02"))

	full2 := time.Date(2010, 12, 06, 0, 0, 0, 0, time.UTC)
	sd2 := FullDate{&full2}
	next2 := sd2.NextAnniversary(referenceDate)
	assert.Equal(t, "2015-12-06", next2.Format("2006-01-02"))
}

func TestNextInDays(t *testing.T) {
	referenceDate := time.Date(2015, 10, 31, 15, 45, 58, 0, time.UTC)
	full := time.Date(2010, 11, 1, 0, 0, 0, 0, time.UTC)
	sd1 := FullDate{&full}
	days := sd1.NextInDays(referenceDate)
	assert.Equal(t, 1, days)
}
