package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func CheckoutController(c echo.Context) error {
	return c.JSON(http.StatusOK, "[WIP] Checkout")
}

func PaymentController(c echo.Context) error {
	return c.JSON(http.StatusOK, "[WIP] Payment")
}
