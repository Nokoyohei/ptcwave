package application

import (
  "net/http"
	"github.com/labstack/echo/v4"
)

func Cardlist(c echo.Context) error {

	return c.String(http.StatusOK, "OK")
}