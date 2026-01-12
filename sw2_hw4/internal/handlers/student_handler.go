package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"sw2_hw3/internal/storage/postgres"
)

func SetupStudentRoutes(e *echo.Echo, storage *postgres.Storage) {
	e.GET("/students/:id", GetStudentByStudentID(storage))
}

func GetStudentByStudentID(storage *postgres.Storage) echo.HandlerFunc {
	return func(c echo.Context) error {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "invalid id")
		}

		student, err := storage.GetStudentByID(context.Background(), id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, student)
	}
}
