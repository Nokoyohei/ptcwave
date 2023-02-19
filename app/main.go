package main

import (
	"fmt"

	"github.com/Nokoyohei/ptcwave/v1/application"
	"github.com/Nokoyohei/ptcwave/v1/database"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", application.Cardlist)

	database.Connect()
	defer database.DB.Close()

	if err := database.DB.Ping(); err != nil {
		fmt.Printf("db connection error \n")
	}

	e.Logger.Fatal(e.Start(":80"))
}
