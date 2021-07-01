package controllers

import (
	"acp9-redy-gigih/config"
	"acp9-redy-gigih/models/category"
	"acp9-redy-gigih/models/product"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func AddProductData() bool {
	prod := product.Product{Name: "Belajar Pemrograman Golang", Price:115000, Stockint: 10, CategoryID: 1}
	err := config.DB.Create(&prod)
	if err != nil {
		return false
	}
	return true
}
func TestGetProductsController(t *testing.T) {
	config.InitDBTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&category.Category{})
	config.DB.Migrator().AutoMigrate(&category.Category{})
	AddCategoryData()
	config.DB.Migrator().DropTable(&product.Product{})
	config.DB.Migrator().AutoMigrate(&product.Product{})
	AddProductData()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/products")
	if assert.NoError(t, GetProductsController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		var responseProducts product.ProductResponse
		fmt.Println(body)
		json.Unmarshal([]byte(body), &responseProducts)

		assert.Equal(t, responseProducts.Status, true)
		assert.Equal(t, len(responseProducts.Data), 1)
		assert.Equal(t, responseProducts.Data[0].Name, "Belajar Pemrograman Golang")
	}
}

func TestFailGetProductsController(t *testing.T) {
	config.InitDBTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&product.Product{})
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/products")
	if assert.NoError(t, GetProductsController(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body := rec.Body.String()
		var responseProducts product.ProductResponse
		fmt.Println(body)
		json.Unmarshal([]byte(body), &responseProducts)

		assert.Equal(t, responseProducts.Status, false)
		assert.Equal(t, len(responseProducts.Data), 0)
	}
}

func TestGetProductsByCategoryController(t *testing.T) {
	config.InitDBTest()
	e := echo.New()
	//models := []interface{}{&product.Product{}, &category.Category{}}
	config.DB.Migrator().DropTable(&category.Category{})
	config.DB.Migrator().AutoMigrate(&category.Category{})
	AddCategoryData()
	config.DB.Migrator().DropTable(&product.Product{})
	config.DB.Migrator().AutoMigrate(&product.Product{})
	AddProductData()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/products/:categorySlug")
	c.SetParamNames("categorySlug")
	c.SetParamValues("buku")
	if assert.NoError(t, GetProductsByCategoryController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		var responseProducts product.ProductResponse
		fmt.Println(body)
		json.Unmarshal([]byte(body), &responseProducts)

		assert.Equal(t, responseProducts.Status, true)
		assert.Equal(t, len(responseProducts.Data), 1)
		assert.Equal(t, responseProducts.Data[0].Name, "Belajar Pemrograman Golang")
	}
}

func TestFailGetCategoryProductsByCategoryController(t *testing.T) {
	config.InitDBTest()
	e := echo.New()
	//models := []interface{}{&product.Product{}, &category.Category{}}
	config.DB.Migrator().DropTable(&category.Category{})
	config.DB.Migrator().DropTable(product.Product{})
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/products/:categorySlug")
	c.SetParamNames("categorySlug")
	c.SetParamValues("buku")
	if assert.NoError(t, GetProductsByCategoryController(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body := rec.Body.String()
		var responseProducts product.ProductResponse
		fmt.Println(body)
		json.Unmarshal([]byte(body), &responseProducts)

		assert.Equal(t, responseProducts.Status, false)
		assert.Equal(t, len(responseProducts.Data), 0)
	}
}

func TestFailGetProductsByCategoryController(t *testing.T) {
	config.InitDBTest()
	e := echo.New()
	//models := []interface{}{&product.Product{}, &category.Category{}}
	config.DB.Migrator().DropTable(&category.Category{})
	config.DB.Migrator().AutoMigrate(&category.Category{})
	AddCategoryData()
	config.DB.Migrator().DropTable(product.Product{})
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/products/:categorySlug")
	c.SetParamNames("categorySlug")
	c.SetParamValues("buku")
	if assert.NoError(t, GetProductsByCategoryController(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body := rec.Body.String()
		var responseProducts product.ProductResponse
		fmt.Println(body)
		json.Unmarshal([]byte(body), &responseProducts)

		assert.Equal(t, responseProducts.Status, false)
		assert.Equal(t, len(responseProducts.Data), 0)
	}
}
