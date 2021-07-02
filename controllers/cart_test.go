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
		ProductID: 1,
		Qty: 2,
	}
	mockDBCartUpdate = models.Cart{
		Qty: 5,
	}
)

func AddCartData() bool {
	cartRow := models.Cart{UserId: 0, ProductID: 1, Qty: 2}
	err := config.DB.Create(&cartRow)
	if err != nil {
		return false
	}
	return true
}

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
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/cart")
	if assert.NoError(t, AddCartController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		var responseCart models.CartResponse
		json.Unmarshal([]byte(body), &responseCart)

		assert.Equal(t, true, responseCart.Status)
		assert.Equal(t, "Add to cart success", responseCart.Message)
	}
}

func TestAddCartControllerSuccessUpdateExists(t *testing.T) {
	config.InitDBTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Cart{})
	config.DB.Migrator().AutoMigrate(&models.Cart{})
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	AddUserData()
	AddCartData()
	token, _ := middlewares.GenerateToken(1)
	body, _ := json.Marshal(mockDBCart)
	r := ioutil.NopCloser(bytes.NewReader(body))
	req := httptest.NewRequest(http.MethodGet, "/", r)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/cart")
	if assert.NoError(t, AddCartController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		var responseCart models.CartResponse
		json.Unmarshal([]byte(body), &responseCart)

		assert.Equal(t, true, responseCart.Status)
		assert.Equal(t, "Add to cart success", responseCart.Message)
	}
}

func TestAddCartControllerFailNoTable(t *testing.T) {
	config.InitDBTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Cart{})
	token, _ := middlewares.GenerateToken(1)
	body, _ := json.Marshal(mockDBCart)
	r := ioutil.NopCloser(bytes.NewReader(body))
	req := httptest.NewRequest(http.MethodGet, "/", r)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/cart")
	if assert.NoError(t, AddCartController(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body := rec.Body.String()
		var responseCart models.CartResponse
		json.Unmarshal([]byte(body), &responseCart)

		assert.Equal(t, false, responseCart.Status)
		assert.Equal(t, "Check same item failed", responseCart.Message)
	}
}

func TestGetCartControllerSuccess(t *testing.T) {
	config.InitDBTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Cart{})
	config.DB.Migrator().AutoMigrate(&models.Cart{})
	AddCartData()
	token, _ := middlewares.GenerateToken(1)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/cart")
	if assert.NoError(t, GetCartController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		var responseCart models.CartResponse
		json.Unmarshal([]byte(body), &responseCart)

		assert.Equal(t, true, responseCart.Status)
		assert.Equal(t, "Get cart success", responseCart.Message)
	}
}

func TestGetCartControllerFailNoTable(t *testing.T) {
	config.InitDBTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Cart{})
	token, _ := middlewares.GenerateToken(1)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/cart")
	if assert.NoError(t, GetCartController(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body := rec.Body.String()
		var responseCart models.CartResponse
		json.Unmarshal([]byte(body), &responseCart)

		assert.Equal(t, false, responseCart.Status)
		assert.Equal(t, "Get cart failed", responseCart.Message)
		assert.Equal(t, 0, len(responseCart.Data))
	}
}

func TestUpdateCartControllerSuccess(t *testing.T) {
	config.InitDBTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Cart{})
	config.DB.Migrator().AutoMigrate(&models.Cart{})
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	AddUserData()
	AddCartData()
	token, _ := middlewares.GenerateToken(1)
	body, _ := json.Marshal(mockDBCartUpdate)
	r := ioutil.NopCloser(bytes.NewReader(body))
	req := httptest.NewRequest(http.MethodGet, "/", r)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/cart")
	c.SetParamNames("productId")
	c.SetParamValues("1")
	fmt.Println(token)
	if assert.NoError(t, UpdateCartController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		var responseCart models.CartResponse
		json.Unmarshal([]byte(body), &responseCart)

		assert.Equal(t, true, responseCart.Status)
		assert.Equal(t, "Update cart success", responseCart.Message)
	}
}

func TestUpdateCartControllerFailNoTable(t *testing.T) {
	config.InitDBTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Cart{})
	token, _ := middlewares.GenerateToken(1)
	body, _ := json.Marshal(mockDBCartUpdate)
	r := ioutil.NopCloser(bytes.NewReader(body))
	req := httptest.NewRequest(http.MethodGet, "/", r)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/cart")
	c.SetParamNames("productId")
	c.SetParamValues("1")
	fmt.Println(token)
	if assert.NoError(t, UpdateCartController(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body := rec.Body.String()
		var responseCart models.CartResponse
		json.Unmarshal([]byte(body), &responseCart)

		assert.Equal(t, false, responseCart.Status)
		assert.Equal(t, "Update cart failed", responseCart.Message)
	}
}

func TestUpdateCartControllerFailRowNotExists(t *testing.T) {
	config.InitDBTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Cart{})
	config.DB.Migrator().AutoMigrate(&models.Cart{})
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	AddUserData()
	AddCartData()
	token, _ := middlewares.GenerateToken(1)
	body, _ := json.Marshal(mockDBCartUpdate)
	r := ioutil.NopCloser(bytes.NewReader(body))
	req := httptest.NewRequest(http.MethodGet, "/", r)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/cart")
	c.SetParamNames("productId")
	c.SetParamValues("4")
	fmt.Println(token)
	if assert.NoError(t, UpdateCartController(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body := rec.Body.String()
		var responseCart models.CartResponse
		json.Unmarshal([]byte(body), &responseCart)

		assert.Equal(t, false, responseCart.Status)
		assert.Equal(t, "Nothing updated", responseCart.Message)
	}
}

func TestDeleteCartControllerSuccess(t *testing.T) {
	config.InitDBTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Cart{})
	config.DB.Migrator().AutoMigrate(&models.Cart{})
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	AddUserData()
	AddCartData()
	token, _ := middlewares.GenerateToken(1)
	body, _ := json.Marshal(mockDBCartUpdate)
	r := ioutil.NopCloser(bytes.NewReader(body))
	req := httptest.NewRequest(http.MethodGet, "/", r)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/cart")
	c.SetParamNames("productId")
	c.SetParamValues("1")
	fmt.Println(token)
	if assert.NoError(t, DeleteCartController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		var responseCart models.CartResponse
		json.Unmarshal([]byte(body), &responseCart)

		assert.Equal(t, true, responseCart.Status)
		assert.Equal(t, "Delete cart success", responseCart.Message)
	}
}

func TestDeleteCartControllerFailNoTable(t *testing.T) {
	config.InitDBTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Cart{})
	token, _ := middlewares.GenerateToken(1)
	body, _ := json.Marshal(mockDBCartUpdate)
	r := ioutil.NopCloser(bytes.NewReader(body))
	req := httptest.NewRequest(http.MethodGet, "/", r)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/cart")
	c.SetParamNames("productId")
	c.SetParamValues("1")
	fmt.Println(token)
	if assert.NoError(t, DeleteCartController(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body := rec.Body.String()
		var responseCart models.CartResponse
		json.Unmarshal([]byte(body), &responseCart)

		assert.Equal(t, false, responseCart.Status)
		assert.Equal(t, "Delete cart failed", responseCart.Message)
	}
}