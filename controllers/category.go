package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func GetCategoriesController(c echo.Context) error {
	return c.JSON(http.StatusOK, "[WIP] Categories")
}

