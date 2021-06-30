package controllers

import (
	"acp9-redy-gigih/config"
	"acp9-redy-gigih/middlewares"
	"acp9-redy-gigih/models/user"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func ourEncrypt(plain string) string {
	bytePlain := []byte(plain)
	hashed, err := bcrypt.GenerateFromPassword(bytePlain, bcrypt.MinCost)
	if err != nil {
		panic(err.Error())
	}
	return string(hashed)
}

func ourCompare(hashed string, plain []byte) bool {
	// Since we'll be getting the hashed password from the DB it will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashed)
	err := bcrypt.CompareHashAndPassword(byteHash, plain)
	if err != nil {
		return err == nil
	}
	return true
}

func RegisterController(e echo.Context) error {
	input := user.User{}
	e.Bind(&input)
	userDB := user.User{}
	userDB.Name = input.Name
	userDB.Email = input.Email
	userDB.Password = ourEncrypt(input.Password)
	err := config.DB.Create(&userDB).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, user.UserResponse{
			false, "Registration failed", nil,
		})
	}
	return e.JSON(http.StatusOK, user.UserResponseSingle{
		true, "Registration success", userDB,
	})
}

func ConfirmationController(e echo.Context) error {
	return e.JSON(http.StatusOK, "[WIP] Register Confirmation")
}

func LoginController(e echo.Context) error {
	input := user.User{}
	e.Bind(&input)
	userDB := user.User{}
	userResponse := user.UserLogin{}
	err := config.DB.Debug().Model(userDB).Where("email = ?", input.Email).Find(&userDB).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, user.UserResponse{
			false, "Wrong email", nil,
		})
	}
	if !ourCompare(userDB.Password, []byte(input.Password)) {
		return e.JSON(http.StatusInternalServerError, user.UserResponse{
			false, "Wrong password", nil,
		})
	}
	userResponse.Token, err = middlewares.GenerateToken(int(userDB.ID))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, user.UserResponse{
			false, "Generata token failed", nil,
		})
	}
	userResponse.ID = userDB.ID
	userResponse.Name = userDB.Name
	userResponse.Email = userDB.Email
	userResponse.Role = userDB.Role
	userResponse.Status = userDB.Status
	return e.JSON(http.StatusOK, user.UserResponseLogin{
		true, "Login success", userResponse,
	})
}
