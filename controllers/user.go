package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterController(e echo.Context) error {
	return e.JSON(http.StatusOK, "[WIP] Register")
}

func ConfirmationController(e echo.Context) error {
	return e.JSON(http.StatusOK, "[WIP] Register Confirmation")
}

func LoginController(e echo.Context) error {
	return e.JSON(http.StatusOK, "[WIP] Login")
}
