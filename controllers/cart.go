package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddCartController(e echo.Context) error {
	return e.JSON(http.StatusOK, "[WIP] Add to Cart")
}

func GetCartController(e echo.Context) error {
	return e.JSON(http.StatusOK, "[WIP] Get Cart")
}

func UpdateCartController(e echo.Context) error {
	return e.JSON(http.StatusOK, "[WIP] Update Cart")
}

func DeleteCartController(e echo.Context) error {
	return e.JSON(http.StatusOK, "[WIP] Delete Cart")
}
