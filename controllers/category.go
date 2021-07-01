package controllers

import (
	"acp9-redy-gigih/config"
	"acp9-redy-gigih/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCategoriesController(e echo.Context) error {
	var categories []models.Category
	err := config.DB.Debug().Model(&models.Category{}).Find(&categories).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.CategoryResponse{
			false, "Failed get database category", nil,
		})
	}
	return e.JSON(http.StatusOK, models.CategoryResponse{
		true, "Success", categories,
	})
}
