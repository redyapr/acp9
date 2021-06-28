package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCategoriesController(c echo.Context) error {
	return c.JSON(http.StatusOK, "[WIP] Categories")
}
