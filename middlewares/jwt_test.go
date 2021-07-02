package middlewares

import (
	"acp9-redy-gigih/config"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func generateTokenTest(userId int) (string, error){
	claims := jwt.MapClaims{
		"authorized": true,
		"userId":     userId,
		"exp":        time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Env("JWT_SECRET")))
}

func TestGenerateToken(t *testing.T) {
	expectation, _ := generateTokenTest(1)
	actual, _ := GenerateToken(1)
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestExtractToken(t *testing.T) {
	token, _ := generateTokenTest(1)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	var expectation float64 = 0
	actual := ExtractToken(c)
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}
