package main

import (
	"fmt"

	"github.com/Nokoyohei/ptcwave/v1/application"
	"github.com/Nokoyohei/ptcwave/v1/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.GET("/", application.Cardlist)
	e.GET("/subprice", application.SubPrice)

	database.Connect()
	defer database.DB.Close()

	if err := database.DB.Ping(); err != nil {
		fmt.Printf("db connection error \n")
	}

	e.Logger.Fatal(e.Start(":8080"))
}
