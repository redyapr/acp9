package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func IndexController(e echo.Context) error {
	return e.String(http.StatusOK, "Welcome to ACP9 Redy Gigih")
}
