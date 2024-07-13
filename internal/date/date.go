package date

import (
	"math/rand"
	"time"
)

func RandDate() time.Time {
	year := 1600 + rand.Intn(400)
	month := time.Month(1 + rand.Intn(12))
	day := 1 + rand.Intn(numDays(month))

	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func numDays(month time.Month) int {
	switch month {
	case time.January, time.March, time.May, time.July, time.August, time.October, time.December:
		return 31
	case time.April, time.June, time.September, time.November:
		return 30
	case time.February:
		// fuck leap years
		return 28
	default:
		panic("unrecognised month")
	}
}
