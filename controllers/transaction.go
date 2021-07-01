package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CheckoutController(e echo.Context) error {
	// userId := middlewares.ExtractToken(e)
	return e.JSON(http.StatusOK, "[WIP] Checkout")
}

func PaymentController(e echo.Context) error {
	// userId := middlewares.ExtractToken(e)
	return e.JSON(http.StatusOK, "[WIP] Payment")
}
