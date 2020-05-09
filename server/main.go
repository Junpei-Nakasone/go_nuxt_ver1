package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type message struct {
	message string
}

func main() {
	e := echo.New()

	e.GET("/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "test")
	})

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "success")
	})

	e.Start(":8080")
}
