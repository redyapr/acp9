package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func AddCartController(c echo.Context) error {
	return c.JSON(http.StatusOK, "[WIP] Add to Cart")
}

func GetCartController(c echo.Context) error {
	return c.JSON(http.StatusOK, "[WIP] Get Cart")
}

func UpdateCartController(c echo.Context) error {
	return c.JSON(http.StatusOK, "[WIP] Update Cart")
}

func DeleteCartController(c echo.Context) error {
	return c.JSON(http.StatusOK, "[WIP] Delete Cart")
}
