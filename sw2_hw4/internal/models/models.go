package models

import (
	"sw2_hw3/internal/utils"
)

// Student and the rest - I used custom types
type Student struct {
	ID          int            `db:"id"`
	FirstName   string         `db:"first_name"`
	LastName    string         `db:"last_name"`
	Gender      string         `db:"gender"`
	BirthDate   utils.DateOnly `db:"birth_date"`
	YearOfStudy int            `db:"year_of_study"`
	GroupId     int            `db:"group_id"`
}

type Schedule struct {
	ID        int            `db:"id"`
	CourseID  int            `db:"course_id"`
	GroupID   int            `db:"group_id"`
	Date      utils.DateOnly `db:"date"`
	StartTime utils.TimeOnly `db:"start_time"`
	EndTime   utils.TimeOnly `db:"end_time"`
}

type Attendance struct {
	ID        int            `db:"id"`
	CourseID  int            `db:"course_id"`
	Date      utils.DateOnly `db:"date"`
	Visited   bool           `db:"visited"`
	StudentID int            `db:"student_id"`
}

// but for bodies I just used string - I genuinely don't know which is better

// AttendanceBody - JSON is different
type AttendanceBody struct {
	CourseID  int    `json:"subject_id"`
	Date      string `json:"visit_day"`
	Visited   bool   `json:"visited"`
	StudentID int    `json:"student_id"`
}

type AttendanceByCourseIDBody struct {
	StudentID int    `json:"student_id"`
	Date      string `json:"visit_day"`
	Visited   bool   `json:"visited"`
}

type AttendanceByStudentIDBody struct {
	CourseID int    `json:"subject_id"`
	Date     string `json:"visit_day"`
	Visited  bool   `json:"visited"`
}
