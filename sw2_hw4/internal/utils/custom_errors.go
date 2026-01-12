package utils

import "errors"

var ErrNoScheduleFound = errors.New("no schedule found to mark this attendance")
var ErrDuplicateAttendance = errors.New("this attendance instance is already present")
var ErrNonExistingStudent = errors.New("there is no existing student by this ID")
var ErrStudentNotEnrolled = errors.New("student is not enrolled in this course")
