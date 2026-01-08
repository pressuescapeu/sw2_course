package main

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"sw2_hw3/internal/storage/postgres"

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

	// TODO: POST /attendance/subject - body {subject_id, visit_day, visited (boolean), student_id}

	e.Logger.Fatal(e.Start(":1323"))
}
