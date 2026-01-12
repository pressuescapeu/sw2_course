package main

import (
	"os"
	"sw2_hw3/internal/handlers"
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

	handlers.SetupStudentRoutes(e, storage)
	handlers.SetupScheduleRoutes(e, storage)
	handlers.SetupAttendanceRoutes(e, storage)

	e.Logger.Fatal(e.Start(":1323"))
}
