package postgres

import (
	"context"
	"sw2_hw3/internal/storage/models"

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
			&schedule.Faculty,
			&schedule.DayOfWeek,
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
			&schedule.Faculty,
			&schedule.DayOfWeek,
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
