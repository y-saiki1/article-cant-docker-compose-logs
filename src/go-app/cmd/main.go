package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	fmt.Println("air doesn't work?")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "bytes_in=${bytes_in},bytes_out=${bytes_out},method=${method}, uri=${uri}, status=${status}, error=${error}\n",
			Output: os.Stdout,
		},
	))
	e.Logger.Fatal(e.Start(":8080"))
}
