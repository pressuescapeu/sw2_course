package postgres

import (
	"context"
	"errors"
	"sw2_hw3/internal/models"
	"sw2_hw3/internal/utils"

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

func (storage *Storage) GetStudentByID(ctx context.Context, id int) (*models.Student, error) {
	const query = `SELECT * FROM students s WHERE s.id = $1;`
	var student models.Student
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
	rows, err := storage.pool.Query(ctx, query)
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

func (storage *Storage) validateAttendanceRequest(ctx context.Context,
	courseID int, visitDay string, studentID int) error {
	visitDate, err := utils.StringDateIntoTimeDate(visitDay)

	if err != nil {
		return err
	}

	const query = `SELECT 
		EXISTS( SELECT id FROM students WHERE id = $1 ) AS student_exists,
		EXISTS( SELECT course_id, date FROM schedule WHERE course_id = $2 AND date = $3 ) AS schedule_exists,
		EXISTS( SELECT student_id, course_id FROM student_courses WHERE course_id = $2 AND student_id = $1 ) AS student_enrolled;
	`

	var studentExists, scheduleExists, studentEnrolled bool
	err = storage.pool.QueryRow(ctx, query, studentID, courseID, visitDate).Scan(&studentExists, &scheduleExists, &studentEnrolled)
	if err != nil {
		return err
	}
	if !studentExists {
		return utils.ErrNonExistingStudent
	}
	if !scheduleExists {
		return utils.ErrNoScheduleFound
	}
	if !studentEnrolled {
		return utils.ErrStudentNotEnrolled
	}
	return nil
}

func (storage *Storage) getDuplicateAttendanceInstance(ctx context.Context,
	courseID int, visitDay string, visited bool, studentID int) ([]models.AttendanceBody, error) {
	visitDate, err := utils.StringDateIntoTimeDate(visitDay)
	if err != nil {
		return nil, err
	}

	const query = `SELECT course_id, TO_CHAR(date, 'DD.MM.YYYY'), visited, student_id FROM attendance WHERE course_id = $1 AND date = $2 AND visited = $3 AND student_id = $4;`
	rows, err := storage.pool.Query(ctx, query, courseID, visitDate, visited, studentID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var attendances []models.AttendanceBody

	for rows.Next() {
		var attendance models.AttendanceBody
		err = rows.Scan(
			&attendance.CourseID,
			&attendance.Date,
			&attendance.Visited,
			&attendance.StudentID,
		)

		if err != nil {
			return nil, err
		}

		attendances = append(attendances, attendance)
	}

	return attendances, rows.Err()
}

func (storage *Storage) MarkAttendance(ctx context.Context,
	courseID int, visitDay string, visited bool, studentID int) error {
	err := storage.validateAttendanceRequest(ctx, courseID, visitDay, studentID)

	if err != nil {
		return err
	}

	attendances, err := storage.getDuplicateAttendanceInstance(ctx, courseID, visitDay, visited, studentID)

	if err != nil {
		return err
	}

	if l := len(attendances); l != 0 {
		return utils.ErrDuplicateAttendance
	}

	const query = `INSERT INTO attendance (course_id, date, visited, student_id) VALUES ($1, $2, $3, $4);`
	visitDate, err := utils.StringDateIntoTimeDate(visitDay)
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
	const query = `SELECT student_id, TO_CHAR(date, 'DD.MM.YYYY'), visited FROM attendance WHERE course_id = $1 ORDER BY date DESC LIMIT 5;`

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
	const query = `SELECT course_id, TO_CHAR(date, 'DD.MM.YYYY'), visited FROM attendance WHERE student_id = $1 ORDER BY date DESC LIMIT 5;`

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
