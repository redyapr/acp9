package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetProductsController(c echo.Context) error {
	return c.JSON(http.StatusOK, "[WIP] Products")
}

func GetProductsByCategoryController(c echo.Context) error {
	return c.JSON(http.StatusOK, "[WIP] Products by Category")
}
