package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// endpoint - "/"
	e.GET("/", func(c echo.Context) error {
		// look for /?name=something
		params := c.QueryParam("name")
		// if nothing then just hello world
		if params != "" {
			// return hello with name in JSON
			res := fmt.Sprintf("Hello, %s!\n", params)
			return c.JSON(http.StatusOK, res)
		}
		return c.JSON(http.StatusBadRequest, "Hello World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
