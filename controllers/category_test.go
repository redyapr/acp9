package controllers

import (
	"acp9-redy-gigih/config"
	"acp9-redy-gigih/models/category"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func AddCategoryData() bool {
	cat := category.Category{Name: "Buku", Description: "Best books on market.", Slug: "buku"}
	err := config.DB.Create(&cat)
	if err != nil {
		return false
	}
	return true
}

func TestGetCategoriesController(t *testing.T) {
	config.InitDBTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&category.Category{})
	config.DB.Migrator().AutoMigrate(&category.Category{})
	AddCategoryData()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/categories")
	if assert.NoError(t, GetCategoriesController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		var responseCategory category.CategoryResponse
		fmt.Println(body)
		json.Unmarshal([]byte(body), &responseCategory)

		assert.Equal(t, responseCategory.Status, true)
		assert.Equal(t, len(responseCategory.Data), 1)
		assert.Equal(t, responseCategory.Data[0].Name, "Buku")
	}
}

func TestFailGetCategoriesController(t *testing.T) {
	config.InitDBTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&category.Category{})
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/categories")
	if assert.NoError(t, GetCategoriesController(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body := rec.Body.String()
		var responseCategory category.CategoryResponse
		fmt.Println(body)
		json.Unmarshal([]byte(body), &responseCategory)

		assert.Equal(t, responseCategory.Status, false)
		assert.Equal(t, len(responseCategory.Data), 0)
	}
}
