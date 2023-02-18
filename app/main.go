package main

import (
	"github.com/Nokoyohei/ptcwave/v1/application"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", application.Cardlist)

	e.Logger.Fatal(e.Start(":80"))
}
