package handlers

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"sw2_hw3/internal/models"
	"sw2_hw3/internal/storage/postgres"
	"sw2_hw3/internal/utils"
	"time"

	"github.com/labstack/echo/v4"
)

func SetupAttendanceRoutes(e *echo.Echo, storage *postgres.Storage) {
	e.POST("/attendance/subject", CreateAttendanceInstance(storage))
	e.GET("/attendanceBySubjectId/:id", GetAllAttendanceByCourseID(storage))
	e.GET("/attendanceByStudentId/:id", GetAllAttendanceByStudentID(storage))
}

func CreateAttendanceInstance(storage *postgres.Storage) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqBody models.AttendanceBody
		if err := c.Bind(&reqBody); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		err := storage.MarkAttendance(context.Background(), reqBody.CourseID, reqBody.Date, reqBody.Visited, reqBody.StudentID)
		if err != nil {
			var timeParseErr *time.ParseError
			if errors.Is(err, utils.ErrNoScheduleFound) || errors.Is(err, utils.ErrDuplicateAttendance) ||
				errors.Is(err, utils.ErrNonExistingStudent) || errors.Is(err, utils.ErrStudentNotEnrolled) {
				return c.JSON(http.StatusBadRequest, err.Error())
			} else if errors.As(err, &timeParseErr) {
				return c.JSON(http.StatusBadRequest, "invalid date format, expected DD.MM.YYYY")
			}
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, "attendance marked successfully")
	}
}

func GetAllAttendanceByCourseID(storage *postgres.Storage) echo.HandlerFunc {
	return func(c echo.Context) error {
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
	}
}

func GetAllAttendanceByStudentID(storage *postgres.Storage) echo.HandlerFunc {
	return func(c echo.Context) error {
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
	}
}
