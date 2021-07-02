package controllers

import (
	"acp9-redy-gigih/config"
	"acp9-redy-gigih/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetProductsController(c echo.Context) error {
	var products []models.Product
	err := config.DB.Debug().Preload("Category").Find(&products).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ProductResponse{
			false, "Failed get database product", nil,
		})
	}
	return c.JSON(http.StatusOK, models.ProductResponse{
		true, "Success", products,
	})
}

func GetProductsByCategoryController(c echo.Context) error {
	categorySlug := c.Param("categorySlug")

	prod := []models.Product{}
	category := models.Category{}
	c.Bind(&prod)
	err := config.DB.Debug().Model(category).Where("slug = ?", categorySlug).Find(&category).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ProductResponse{
			false, "categorySlug not found", nil,
		})
	}

	err = config.DB.Debug().Preload("Category").Where("category_id = ?", category.ID).Find(&prod).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ProductResponse{
			false, "products not found", nil,
		})
	}

	return c.JSON(http.StatusOK, models.ProductResponse{
		true, "Success", prod,
	})
}
