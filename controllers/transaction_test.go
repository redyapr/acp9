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
	mockDBTransaction = models.Transaction{
		ID: 1,
		UserID: 0,
		Status: "Unpaid",
	}
)

func AddTransactionData() bool {
	transactionRow := models.Transaction{UserID: 0, Status: "Unpaid"}
	err := config.DB.Create(&transactionRow)
	if err != nil {
		return false
	}
	return true
}

func AddTransactionDataPayment() bool {
	transactionRow := models.Transaction{UserID: 5, Status: "Unpaid"}
	err := config.DB.Create(&transactionRow)
	if err != nil {
		return false
	}
	return true
}

func TestGetTransactionsControllerSuccess(t *testing.T) {
	config.InitDBTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Transaction{})
	config.DB.Migrator().AutoMigrate(&models.Transaction{})
	AddTransactionData()
	token, _ := middlewares.GenerateToken(1)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/transactions")
	if assert.NoError(t, GetTransactionsController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		var responseCart models.CartResponse
		json.Unmarshal([]byte(body), &responseCart)

		assert.Equal(t, true, responseCart.Status)
		assert.Equal(t, "Get cart success", responseCart.Message)
	}
}

func TestGetTransactionsControllerFailNoTable(t *testing.T) {
	config.InitDBTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Transaction{})
	token, _ := middlewares.GenerateToken(1)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/transactions")
	if assert.NoError(t, GetTransactionsController(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body := rec.Body.String()
		var responseTransaction models.TransactionResponse
		json.Unmarshal([]byte(body), &responseTransaction)

		assert.Equal(t, false, responseTransaction.Status)
		assert.Equal(t, "Get transactions failed", responseTransaction.Message)
	}
}

func TestCheckoutControllerSuccess(t *testing.T) {
	config.InitDBTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Cart{})
	config.DB.Migrator().AutoMigrate(&models.Cart{})
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	AddUserData()
	AddCartData()
	config.DB.Migrator().DropTable(&models.Transaction{})
	config.DB.Migrator().AutoMigrate(&models.Transaction{})
	config.DB.Migrator().DropTable(&models.TransactionDetail{})
	config.DB.Migrator().AutoMigrate(&models.TransactionDetail{})
	token, _ := middlewares.GenerateToken(1)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/checkout")
	fmt.Println(token)
	if assert.NoError(t, CheckoutController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		var responseTransaction models.TransactionResponse
		json.Unmarshal([]byte(body), &responseTransaction)

		assert.Equal(t, true, responseTransaction.Status)
		assert.Equal(t, "Checkout success", responseTransaction.Message)
	}
}

func TestCheckoutControllerFailNoTableCart(t *testing.T) {
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
	c.SetPath("/checkout")
	if assert.NoError(t, CheckoutController(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body := rec.Body.String()
		var responseTransaction models.TransactionResponse
		json.Unmarshal([]byte(body), &responseTransaction)

		assert.Equal(t, false, responseTransaction.Status)
		assert.Equal(t, "Get cart failed", responseTransaction.Message)
	}
}

func TestCheckoutControllerFailCartNotExists(t *testing.T) {
	config.InitDBTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Cart{})
	config.DB.Migrator().AutoMigrate(&models.Cart{})
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	AddUserData()
	token, _ := middlewares.GenerateToken(1)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/checkout")
	fmt.Println(token)
	if assert.NoError(t, CheckoutController(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body := rec.Body.String()
		var responseTransaction models.TransactionResponse
		json.Unmarshal([]byte(body), &responseTransaction)

		assert.Equal(t, false, responseTransaction.Status)
		assert.Equal(t, "Cart is empty", responseTransaction.Message)
	}
}

func TestCheckoutControllerFailNoTableTransaction(t *testing.T) {
	config.InitDBTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Cart{})
	config.DB.Migrator().AutoMigrate(&models.Cart{})
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	AddUserData()
	AddCartData()
	config.DB.Migrator().DropTable(&models.Transaction{})
	token, _ := middlewares.GenerateToken(1)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/checkout")
	fmt.Println(token)
	if assert.NoError(t, CheckoutController(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body := rec.Body.String()
		var responseTransaction models.TransactionResponse
		json.Unmarshal([]byte(body), &responseTransaction)

		assert.Equal(t, false, responseTransaction.Status)
		assert.Equal(t, "Create transaction head failed", responseTransaction.Message)
	}
}

func TestCheckoutControllerFailNoTableDetail(t *testing.T) {
	config.InitDBTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Cart{})
	config.DB.Migrator().AutoMigrate(&models.Cart{})
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	AddUserData()
	AddCartData()
	config.DB.Migrator().DropTable(&models.Transaction{})
	config.DB.Migrator().AutoMigrate(&models.Transaction{})
	config.DB.Migrator().DropTable(&models.TransactionDetail{})
	token, _ := middlewares.GenerateToken(1)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/checkout")
	fmt.Println(token)
	if assert.NoError(t, CheckoutController(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body := rec.Body.String()
		var responseTransaction models.TransactionResponse
		json.Unmarshal([]byte(body), &responseTransaction)

		assert.Equal(t, false, responseTransaction.Status)
		assert.Equal(t, "Create transaction detail failed", responseTransaction.Message)
	}
}

func TestPaymentControllerSuccess(t *testing.T) {
	config.InitDBTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Transaction{})
	config.DB.Migrator().AutoMigrate(&models.Transaction{})
	AddTransactionData()
	token, _ := middlewares.GenerateToken(1)
	body, _ := json.Marshal(mockDBTransaction)
	r := ioutil.NopCloser(bytes.NewReader(body))
	req := httptest.NewRequest(http.MethodGet, "/", r)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/payment")
	fmt.Println(token)
	if assert.NoError(t, PaymentController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		var responseTransaction models.TransactionResponse
		json.Unmarshal([]byte(body), &responseTransaction)

		assert.Equal(t, true, responseTransaction.Status)
		assert.Equal(t, "Payment success", responseTransaction.Message)
	}
}

func TestPaymentControllerFailNoTable(t *testing.T) {
	config.InitDBTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Transaction{})
	token, _ := middlewares.GenerateToken(1)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/payment")
	fmt.Println(token)
	if assert.NoError(t, PaymentController(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body := rec.Body.String()
		var responseTransaction models.TransactionResponse
		json.Unmarshal([]byte(body), &responseTransaction)

		assert.Equal(t, false, responseTransaction.Status)
		assert.Equal(t, "Payment failed", responseTransaction.Message)
	}
}

func TestPaymentControllerFailRowNotExists(t *testing.T) {
	config.InitDBTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.Transaction{})
	config.DB.Migrator().AutoMigrate(&models.Transaction{})
	AddTransactionDataPayment()
	token, _ := middlewares.GenerateToken(1)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/payment")
	fmt.Println(token)
	if assert.NoError(t, PaymentController(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body := rec.Body.String()
		var responseTransaction models.TransactionResponse
		json.Unmarshal([]byte(body), &responseTransaction)

		assert.Equal(t, false, responseTransaction.Status)
		assert.Equal(t, "Transaction ID not found", responseTransaction.Message)
	}
}