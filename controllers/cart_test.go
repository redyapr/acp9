package controllers

import (
	"acp9-redy-gigih/config"
	"acp9-redy-gigih/middlewares"
	"acp9-redy-gigih/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	mockDBCart = models.Cart{
		UserId: 1,
		ProductID: 1,
		Qty: 2,
	}
)

func TestAddCartControllerSuccess(t *testing.T) {
	config.InitDBTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Cart{})
	config.DB.Migrator().AutoMigrate(&models.Cart{})
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	AddUserData()
	token, _ := middlewares.GenerateToken(1)
	body, _ := json.Marshal(mockDBCart)
	r := ioutil.NopCloser(bytes.NewReader(body))
	req := httptest.NewRequest(http.MethodGet, "/", r)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/cart")
	fmt.Println(c)
	fmt.Println(token)
	if assert.NoError(t, AddCartController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		var responseCart models.CartResponse
		fmt.Println(body)
		json.Unmarshal([]byte(body), &responseCart)

		assert.Equal(t, true, responseCart.Status)
		assert.Equal(t, "Add to cart success", responseCart.Message)
	}
}
