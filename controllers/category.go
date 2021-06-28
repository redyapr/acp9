package controllers

import (
	"acp9-redy-gigih/config"
	"acp9-redy-gigih/models/category"
	"github.com/labstack/echo"
	"net/http"
)

func GetCategoriesController(c echo.Context) error {
	var categories []category.Category

	err := config.DB.Debug().Model(&category.Category{}).Find(&categories).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, category.CategoryResponse{
			false,"Failed get database category", nil,
		})
	}
	return c.JSON(http.StatusOK, category.CategoryResponse{
		true, "Success", categories,
	})
}