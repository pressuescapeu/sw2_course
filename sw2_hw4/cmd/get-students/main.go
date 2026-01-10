package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strconv"
	"sw2_hw3/internal/models"
	"sw2_hw3/internal/storage/postgres"
	"sw2_hw3/internal/utils"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// loading the .env file in the project root
	godotenv.Load()

	e := echo.New()

	// connect to the database
	connString := os.Getenv("DATABASE_URL")
	if connString == "" {
		panic("DATABASE_URL not set")
	}

	storage, err := postgres.NewConnection(connString)
	if err != nil {
		panic(err)
	}

	e.GET("/students/:id", func(c echo.Context) error {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "invalid id")
		}
		// context.Background - context that never cancels, has no deadline
		student, err := storage.GetStudentByID(context.Background(), id)

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error()) // to show the full error
		}

		return c.JSON(http.StatusOK, student)
	})

	e.GET("/all_class_schedule", func(c echo.Context) error {
		schedules, err := storage.GetScheduleForAll(context.Background())

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, schedules)
	})

	e.GET("/schedule/group/:id", func(c echo.Context) error {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)

		if err != nil {
			return c.JSON(http.StatusBadRequest, "invalid id")
		}

		schedules, err := storage.GetScheduleByGroupID(context.Background(), id)

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, schedules)
	})

	e.POST("/attendance/subject", func(c echo.Context) error {
		var reqBody models.AttendanceBody
		// Bind will unmarshall it into the interface of AttendanceBody
		if err := c.Bind(&reqBody); err != nil {
			// return a 400 with error
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		// here we mark attendance - Date will be string, and converted later in postgres.go
		err := storage.MarkAttendance(context.Background(), reqBody.CourseID, reqBody.Date, reqBody.Visited, reqBody.StudentID)
		if err != nil {
			// type of error in case user puts in invalid date format
			var timeParseErr *time.ParseError
			if err == utils.ErrNoScheduleFound {
				// 400 bc user put in wrong info about attendance
				return c.JSON(http.StatusBadRequest, err.Error())
			} else if errors.As(err, &timeParseErr) {
				// 400 bc user put in invalid date and then utils couldn't parse
				return c.JSON(http.StatusBadRequest, "invalid date format, expected DD.MM.YYYY")
			}
			// return 500 with error
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		// return 201 with a success message
		return c.JSON(http.StatusCreated, "attendance marked successfully")
	})

	e.GET("/attendanceBySubjectId/:id", func(c echo.Context) error {
		id := c.Param("id")
		courseID, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "invalid id")
		}
		attendances, err := storage.GetAttendanceByCourseID(context.Background(), courseID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, attendances)
	})

	e.GET("/attendanceByStudentId/:id", func(c echo.Context) error {
		id := c.Param("id")
		studentID, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "invalid id")
		}
		attendances, err := storage.GetAttendanceByStudentID(context.Background(), studentID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, attendances)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
