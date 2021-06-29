package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCategoriesController(e echo.Context) error {
	return e.JSON(http.StatusOK, "[WIP] Categories")
}
