package main

import (
	"fmt"
	"os"

	"github.com/Nokoyohei/ptcwave/v1/application"
	"github.com/Nokoyohei/ptcwave/v1/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	e := echo.New()

	env := os.Getenv("ENV")
	if env == "prod" {
		e.AutoTLSManager.HostPolicy = autocert.HostWhitelist(os.Getenv("HOST"))
		e.AutoTLSManager.Cache = autocert.DirCache("/var/www/certs")
		e.Pre(middleware.HTTPSWWWRedirect())
	}

	e.Use(middleware.CORS())
	e.GET("/", application.Cardlist)
	e.GET("/subprice", application.SubPrice)

	database.Connect()
	defer database.DB.Close()

	if err := database.DB.Ping(); err != nil {
		fmt.Printf("db connection error \n")
	}

	switch env {
	case "prod":
		e.Logger.Fatal(e.StartAutoTLS(":443"))
	case "dev":
		defaultAddr := ":80"
		e.Logger.Fatal(e.Start(defaultAddr))
	default:
		defaultAddr := ":8080"
		e.Logger.Fatal(e.Start(defaultAddr))
	}
}
