package controllers

import (
	"acp9-redy-gigih/config"
	"acp9-redy-gigih/models/product"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetProductsController(c echo.Context) error {
	var products []product.Product

	err := config.DB.Debug().Model(&product.Product{}).Find(&products).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, product.ProductResponse{
			false,"Failed get database product", nil,
		})
	}
	return c.JSON(http.StatusOK, product.ProductResponse{
		true, "Success", products,
	})
}

func GetProductsByCategoryController(c echo.Context) error {
	return c.JSON(http.StatusOK, "[WIP] Products by Category")
}
