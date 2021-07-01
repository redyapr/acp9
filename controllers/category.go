package controllers

import (
	"acp9-redy-gigih/config"
	"acp9-redy-gigih/models/category"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCategoriesController(e echo.Context) error {
	var categories []category.Category
	err := config.DB.Debug().Model(&category.Category{}).Find(&categories).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, category.CategoryResponse{
			false, "Failed get database category", nil,
		})
	}
	return e.JSON(http.StatusOK, category.CategoryResponse{
		true, "Success", categories,
	})
}
