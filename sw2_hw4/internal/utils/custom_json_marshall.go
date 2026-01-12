package utils

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type DateOnly time.Time

// Scan - convert db value into the custom type
func (d *DateOnly) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	if t, ok := value.(time.Time); ok {
		*d = DateOnly(t)
		return nil
	}

	if str, ok := value.(string); ok {
		parsed, err := time.Parse("2006-01-02", str)
		if err != nil {
			return fmt.Errorf("cannot parse date string %s: %w", str, err)
		}
		*d = DateOnly(parsed)
		return nil
	}
	return fmt.Errorf("cannot scan %T into DateOnly", value)
}

func (d DateOnly) Value() (driver.Value, error) {
	return time.Time(d), nil
}

func (d DateOnly) MarshalJSON() ([]byte, error) {
	output := time.Time(d).Format((`"02.01.2006"`))
	return []byte(output), nil
}

type TimeOnly time.Time

func (t *TimeOnly) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	if tm, ok := value.(time.Time); ok {
		*t = TimeOnly(tm)
		return nil
	}

	if str, ok := value.(string); ok {
		parsed, err := time.Parse("15:04:05", str)
		if err != nil {
			return fmt.Errorf("cannot parse time string %s: %w", str, err)
		}
		*t = TimeOnly(parsed)
		return nil
	}
	return fmt.Errorf("cannot scan %T into TimeOnly", value)
}

func (t TimeOnly) Value() (driver.Value, error) {
	return time.Time(t), nil
}

func (t TimeOnly) MarshalJSON() ([]byte, error) {
	output := time.Time(t).Format(`"15:04"`)
	return []byte(output), nil
}
