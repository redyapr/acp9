package controllers

import (
	"acp9-redy-gigih/config"
	"acp9-redy-gigih/models/category"
	"acp9-redy-gigih/models/product"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetProductsController(c echo.Context) error {
	var products []product.Product

	err := config.DB.Debug().Preload("Category").Find(&products).Error
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
	categorySlug := c.Param("categorySlug")

	prod := []product.Product{}
	category := category.Category{}
	c.Bind(&prod)
	err := config.DB.Debug().Model(category).Where("slug = ?", categorySlug).Find(&category).Error
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, product.ProductResponse{
			false, "categorySlug not found", nil,
		})
	}

	err = config.DB.Debug().Preload("Category").Where("category_id = ?", category.ID).Find(&prod).Error
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, product.ProductResponse{
			false, "products not found", nil,
		})
	}

	return c.JSON(http.StatusOK, product.ProductResponse{
		true, "Success", prod,
	})
}
