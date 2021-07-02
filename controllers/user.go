package controllers

import (
	"acp9-redy-gigih/config"
	"acp9-redy-gigih/middlewares"
	"acp9-redy-gigih/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func ourEncrypt(plain string) string {
	bytePlain := []byte(plain)
	hashed, _ := bcrypt.GenerateFromPassword(bytePlain, bcrypt.MinCost)
	return string(hashed)
}

func ourCompare(hashed string, plain []byte) bool {
	byteHash := []byte(hashed)
	err := bcrypt.CompareHashAndPassword(byteHash, plain)
	if err != nil {
		return err == nil
	}
	return true
}

func RegisterController(e echo.Context) error {
	input := models.User{}
	e.Bind(&input)
	if input.Name == "" || input.Email == "" || input.Password == "" {
		return e.JSON(http.StatusInternalServerError, models.UserResponse{
			false, "Name/Email/Password cannot empty", nil,
		})
	}
	userDB := models.User{}
	userDB.Name = input.Name
	userDB.Email = input.Email
	userDB.Password = ourEncrypt(input.Password)
	err := config.DB.Create(&userDB).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.UserResponse{
			false, "Registration failed", nil,
		})
	}
	return e.JSON(http.StatusOK, models.UserResponseSingle{
		true, "Registration success", userDB,
	})
}

func LoginController(e echo.Context) error {
	input := models.User{}
	e.Bind(&input)
	userDB := models.User{}
	userResponse := models.UserLogin{}
	err := config.DB.Debug().Model(userDB).Where("email = ?", input.Email).Find(&userDB).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.UserResponse{
			false, "Database error", nil,
		})
	}
	count := config.DB.Debug().Model(userDB).Where("email = ?", input.Email).Find(&userDB).RowsAffected
	if count == 0 {
		return e.JSON(http.StatusInternalServerError, models.UserResponse{
			false, "Wrong email", nil,
		})
	}
	if !ourCompare(userDB.Password, []byte(input.Password)) {
		return e.JSON(http.StatusInternalServerError, models.UserResponse{
			false, "Wrong password", nil,
		})
	}
	userResponse.Token, _ = middlewares.GenerateToken(int(userDB.ID))
	userResponse.ID = userDB.ID
	userResponse.Name = userDB.Name
	userResponse.Email = userDB.Email
	userResponse.Role = userDB.Role
	userResponse.Status = userDB.Status
	return e.JSON(http.StatusOK, models.UserResponseLogin{
		true, "Login success", userResponse,
	})
}
