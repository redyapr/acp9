package controllers

import (
	"acp9-redy-gigih/config"
	"acp9-redy-gigih/models/category"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetCategoriesController(t *testing.T) {
	e := config.SetupEchoDB()
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
		assert.Equal(t, len(responseCategory.Data), 2)
		assert.Equal(t, responseCategory.Data[0].Name, "Buku")
	}

}
