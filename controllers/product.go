package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func GetProductsController(c echo.Context) error {
	return c.JSON(http.StatusOK, "[WIP] Products")
}

func GetProductsByCategoryController(c echo.Context) error {
	return c.JSON(http.StatusOK, "[WIP] Products by Category")
}