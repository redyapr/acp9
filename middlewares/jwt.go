package middlewares

import (
	"acp9-redy-gigih/config"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func GenerateToken(userId int) (string, error) {
	claims := jwt.MapClaims{
		"authorized": true,
		"userId":     userId,
		"exp":        time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Env("JWT_SECRET")))
}

func ExtractToken(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		return claims["userId"].(int)
	}
	return 0
}
