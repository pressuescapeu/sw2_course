package models

import "time"

type Student struct {
	ID          int       `db:"id"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	Gender      string    `db:"gender"`
	BirthDate   time.Time `db:"birth_date"`
	YearOfStudy int       `db:"year_of_study"`
	GroupId     int       `db:"group_id"`
}

type Schedule struct {
	ID        int       `db:"id"`
	CourseID  int       `db:"course_id"`
	GroupID   int       `db:"group_id"`
	Faculty   string    `db:"faculty"`
	DayOfWeek string    `db:"day_of_week"`
	StartTime time.Time `db:"start_time"`
	EndTime   time.Time `db:"end_time"`
}
