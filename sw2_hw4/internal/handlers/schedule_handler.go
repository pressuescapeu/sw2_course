package handlers

import (
	"context"
	"net/http"
	"strconv"
	"sw2_hw3/internal/storage/postgres"

	"github.com/labstack/echo/v4"
)

func SetupScheduleRoutes(e *echo.Echo, storage *postgres.Storage) {
	e.GET("/all_class_schedule", GetAllClassSchedule(storage))
	e.GET("/schedule/group/:id", GetGroupScheduleByGroupID(storage))
}

func GetAllClassSchedule(storage *postgres.Storage) echo.HandlerFunc {
	return func(c echo.Context) error {
		schedules, err := storage.GetScheduleForAll(context.Background())

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, schedules)
	}
}

func GetGroupScheduleByGroupID(storage *postgres.Storage) echo.HandlerFunc {
	return func(c echo.Context) error {
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
	}
}
