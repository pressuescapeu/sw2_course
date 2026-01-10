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
	Date      time.Time `db:"date"`
	StartTime time.Time `db:"start_time"`
	EndTime   time.Time `db:"end_time"`
}

type Attendance struct {
	ID        int       `db:"id"`
	CourseID  int       `db:"course_id"`
	Date      time.Time `db:"date"`
	Visited   bool      `db:"visited"`
	StudentID int       `db:"student_id"`
}

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
