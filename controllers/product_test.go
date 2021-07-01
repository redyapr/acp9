package controllers

import (
	"acp9-redy-gigih/config"
	"acp9-redy-gigih/models/product"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetProductsController(t *testing.T) {
	e := config.SetupEchoDB()
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
		assert.Equal(t, len(responseProducts.Data), 7)
		assert.Equal(t, responseProducts.Data[0].Name, "Pemrograman Web dengan Node.js dan Javascript")
	}
}

func TestGetProductsByCategoryController(t *testing.T) {
	e := config.SetupEchoDB()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/products/:categorySlug")
	c.SetParamNames("categorySlug")
	c.SetParamValues("barang-elektronik")
	if assert.NoError(t, GetProductsByCategoryController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		var responseProducts product.ProductResponse
		fmt.Println(body)
		json.Unmarshal([]byte(body), &responseProducts)

		assert.Equal(t, responseProducts.Status, true)
		assert.Equal(t, len(responseProducts.Data), 4)
		assert.Equal(t, responseProducts.Data[0].Name, "Speaker Logitech")
	}
}
