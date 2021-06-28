package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CheckoutController(c echo.Context) error {
	return c.JSON(http.StatusOK, "[WIP] Checkout")
}

func PaymentController(c echo.Context) error {
	return c.JSON(http.StatusOK, "[WIP] Payment")
}
