package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetProductsController(e echo.Context) error {
	return e.JSON(http.StatusOK, "[WIP] Products")
}

func GetProductsByCategoryController(e echo.Context) error {
	return e.JSON(http.StatusOK, "[WIP] Products by Category")
}
