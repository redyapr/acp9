package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func IndexController(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to ACP9 Redy Gigih")
}

func RegisterController(c echo.Context) error {
	return c.JSON(http.StatusOK, "[WIP] Register")
}

func ConfirmationController(c echo.Context) error {
	return c.JSON(http.StatusOK, "[WIP] Register Confirmation")
}

func LoginController(c echo.Context) error {
	return c.JSON(http.StatusOK, "[WIP] Login")
}
