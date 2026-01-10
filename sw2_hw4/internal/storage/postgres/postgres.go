package postgres

import (
	"context"
	"errors"
	"sw2_hw3/internal/models"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Storage is used so that it's easier to interact with db
type Storage struct {
	pool *pgxpool.Pool
}

// NewConnection returns *Storage so the pool is shared
func NewConnection(connString string) (*Storage, error) {
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, err
	}

	return &Storage{
		pool: pool,
	}, nil
}

// GetStudentByID uses ctx to manage request lifecycle, uses Background in main
func (storage *Storage) GetStudentByID(ctx context.Context, id int) (*models.Student, error) {
	const query = `
		SELECT * FROM students s WHERE s.id = $1;
	`
	var student models.Student
	// QueryRow returns only one row so we scan right away
	err := storage.pool.QueryRow(ctx, query, id).Scan(
		&student.ID,
		&student.FirstName,
		&student.LastName,
		&student.Gender,
		&student.BirthDate,
		&student.YearOfStudy,
		&student.GroupId,
	)

	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (storage *Storage) GetScheduleForAll(ctx context.Context) ([]models.Schedule, error) {
	const query = `SELECT * FROM schedule;`
	// Query() returns multiple rows so we save them as rows
	rows, err := storage.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var schedules []models.Schedule
	// then we parse these rows one by one and add to schedules slice
	for rows.Next() {
		var schedule models.Schedule
		err = rows.Scan(
			&schedule.ID,
			&schedule.CourseID,
			&schedule.GroupID,
			&schedule.Date,
			&schedule.StartTime,
			&schedule.EndTime,
		)

		if err != nil {
			return nil, err
		}

		schedules = append(schedules, schedule)
	}

	return schedules, rows.Err()
}

func (storage *Storage) GetScheduleByGroupID(ctx context.Context, id int) ([]models.Schedule, error) {
	const query = `SELECT * FROM schedule WHERE group_id = $1;`

	rows, err := storage.pool.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var schedules []models.Schedule

	for rows.Next() {
		var schedule models.Schedule
		err = rows.Scan(
			&schedule.ID,
			&schedule.CourseID,
			&schedule.GroupID,
			&schedule.Date,
			&schedule.StartTime,
			&schedule.EndTime,
		)

		if err != nil {
			return nil, err
		}

		schedules = append(schedules, schedule)
	}

	return schedules, rows.Err()
}

func stringDateIntoTimeDate(input string) (time.Time, error) {
	// idk some go magic date stuff - 2nd jan 2006
	return time.Parse("02.01.2006", input)
}

// GetScheduleByDateAndCourseID - helper function for marking attendance
func (storage *Storage) getScheduleByDateAndCourseID(ctx context.Context,
	courseID int, visitDay string) ([]models.Schedule, error) {
	visitDate, err := stringDateIntoTimeDate(visitDay)
	if err != nil {
		return nil, err
	}
	// first check if there even is a schedule instance like this
	const query = `SELECT * FROM schedule WHERE course_id = $1 AND date = $2;`

	rows, err := storage.pool.Query(ctx, query, courseID, visitDate)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var schedules []models.Schedule

	for rows.Next() {
		var schedule models.Schedule
		err = rows.Scan(
			&schedule.ID,
			&schedule.CourseID,
			&schedule.GroupID,
			&schedule.Date,
			&schedule.StartTime,
			&schedule.EndTime,
		)

		if err != nil {
			return nil, err
		}

		schedules = append(schedules, schedule)
	}

	return schedules, rows.Err()
}

func (storage *Storage) MarkAttendance(ctx context.Context,
	courseID int, visitDay string, visited bool, studentID int) error {
	schedules, err := storage.getScheduleByDateAndCourseID(ctx, courseID, visitDay)

	if err != nil {
		return err
	}
	// also check if schedules is just empty
	if l := len(schedules); l == 0 {
		return errors.New("no schedule found to mark this attendance")
	}

	const query = `INSERT INTO attendance (course_id, date, visited, student_id) VALUES ($1, $2, $3, $4);`
	visitDate, err := stringDateIntoTimeDate(visitDay)
	if err != nil {
		return err
	}

	cmdTag, err := storage.pool.Exec(ctx, query, courseID, visitDate, visited, studentID)

	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return errors.New("no rows were inserted")
	}

	return nil
}

func (storage *Storage) GetAttendanceByCourseID(ctx context.Context,
	id int) ([]models.AttendanceByCourseIDBody, error) {
	const query = `SELECT student_id, date, visited FROM attendance WHERE course_id = $1;`

	rows, err := storage.pool.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var attendanceUnits []models.AttendanceByCourseIDBody

	for rows.Next() {
		var attendanceUnit models.AttendanceByCourseIDBody
		err = rows.Scan(
			&attendanceUnit.StudentID,
			&attendanceUnit.Date,
			&attendanceUnit.Visited,
		)

		if err != nil {
			return nil, err
		}
		attendanceUnits = append(attendanceUnits, attendanceUnit)
	}

	return attendanceUnits, nil
}

// GetAttendanceByStudentID - almost identical to the prev one
func (storage *Storage) GetAttendanceByStudentID(ctx context.Context,
	id int) ([]models.AttendanceByStudentIDBody, error) {
	const query = `SELECT course_id, date, visited FROM attendance WHERE student_id = $1;`

	rows, err := storage.pool.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var attendanceUnits []models.AttendanceByStudentIDBody

	for rows.Next() {
		var attendanceUnit models.AttendanceByStudentIDBody
		err = rows.Scan(
			&attendanceUnit.CourseID,
			&attendanceUnit.Date,
			&attendanceUnit.Visited,
		)

		if err != nil {
			return nil, err
		}
		attendanceUnits = append(attendanceUnits, attendanceUnit)
	}

	return attendanceUnits, nil
}
