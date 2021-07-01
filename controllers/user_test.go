package controllers

import (
	"acp9-redy-gigih/config"
	"acp9-redy-gigih/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	mockDB = models.User{
		Name:     "Gigih",
		Email:    "francinogigih@gmail.com",
		Password: "123",
	}
	mockDBLoginWrongEmail = models.User{
		Email:    "francino@gmail.com",
		Password: "123",
	}
	mockDBLoginWrongPassword = models.User{
		Email:    "francinogigih@gmail.com",
		Password: "12345",
	}
)

func AddUserData() bool {
	userData := models.User{Name: "Gigih", Email: "francinogigih@gmail.com", Password: "$2b$10$q1uk7lDvHd1M7BztlgrU9.t3dg32I/11Yv3uDK0Kcycp4v3BTabi6"}
	err := config.DB.Debug().Create(&userData)
	if err != nil {
		return false
	}
	return true
}

func TestRegisterControllerSuccess(t *testing.T) {
	config.InitDBTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	body, _ := json.Marshal(mockDB)
	r := ioutil.NopCloser(bytes.NewReader(body))
	req := httptest.NewRequest(http.MethodGet, "/", r)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/register")
	if assert.NoError(t, RegisterController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		var responseUser models.UserResponse
		fmt.Println(body)
		json.Unmarshal([]byte(body), &responseUser)

		assert.Equal(t, true, responseUser.Status)
		assert.Equal(t, "Registration success", responseUser.Message)
	}
}

func TestRegisterControllerFailEmptyBody(t *testing.T) {
	config.InitDBTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/register")
	if assert.NoError(t, RegisterController(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body := rec.Body.String()
		var responseUser models.UserResponse
		fmt.Println(body)
		json.Unmarshal([]byte(body), &responseUser)

		assert.Equal(t, false, responseUser.Status)
		assert.Equal(t, "Name/Email/Password cannot empty", responseUser.Message)
	}
}

func TestRegisterControllerFailNoTable(t *testing.T) {
	config.InitDBTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.User{})
	body, _ := json.Marshal(mockDB)
	r := ioutil.NopCloser(bytes.NewReader(body))
	req := httptest.NewRequest(http.MethodGet, "/", r)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/register")
	if assert.NoError(t, RegisterController(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body := rec.Body.String()
		var responseUser models.UserResponse
		fmt.Println(body)
		json.Unmarshal([]byte(body), &responseUser)

		assert.Equal(t, false, responseUser.Status)
		assert.Equal(t, "Registration failed", responseUser.Message)
	}
}

func TestLoginControllerFailWrongPassword(t *testing.T) {
	config.InitDBTest()
	e := echo.New()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	AddUserData()
	body, _ := json.Marshal(&mockDBLoginWrongPassword)
	r := ioutil.NopCloser(bytes.NewReader(body))
	req := httptest.NewRequest(http.MethodGet, "/", r)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/login")
	if assert.NoError(t, LoginController(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body := rec.Body.String()
		var responseUser models.UserResponse
		fmt.Println(body)
		json.Unmarshal([]byte(body), &responseUser)

		assert.Equal(t, false, responseUser.Status)
		assert.Equal(t, "Wrong password", responseUser.Message)
	}
}
