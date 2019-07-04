package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/tech", func(context echo.Context) error {
		return context.JSON(http.StatusOK, "ok")
	})

	e.Logger.Fatal(e.Start(":1234"))
}
