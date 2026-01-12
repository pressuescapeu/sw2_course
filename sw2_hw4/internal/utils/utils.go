package utils

import "time"

func StringDateIntoTimeDate(input string) (time.Time, error) {
	return time.Parse("02.01.2006", input)
}
