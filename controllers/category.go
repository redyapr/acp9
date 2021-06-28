package controllers

import (
	"acp9-redy-gigih/config"
	"acp9-redy-gigih/models/category"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func GetCategoriesController(c echo.Context) error {
	var categories []category.Category

	err := config.DB.Debug().Model(&category.Category{}).Find(&categories).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, category.CategoryResponse{
			false,"Failed get database category", nil,
		})
	}
	//categoriesResponse := category.CategoriesResponse{categories.ID, categories}
	return c.JSON(http.StatusOK, category.CategoryResponse{
		true, "Success", categories,
	})
}

func GetCategoryController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, category.CategoryResponse{
			false, err.Error(), nil,
		})
	}
	cat := category.Category{}
	c.Bind(&cat)
	err = config.DB.Debug().Model(&cat).Where("id = ?", id).Take(&cat).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, category.CategoryResponse{
			false, "category_id not found", nil,
		})
	}
	return c.JSON(http.StatusOK, category.CategoryResponseSingle{
		true,"Success", cat,
	})
}

