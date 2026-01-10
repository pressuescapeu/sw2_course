package utils

import "time"

func StringDateIntoTimeDate(input string) (time.Time, error) {
	// idk some go magic date stuff - 2nd jan 2006
	return time.Parse("02.01.2006", input)
}
