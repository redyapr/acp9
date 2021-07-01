package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	Name      string         `json:"name"`
	Email     string         `gorm:"UNIQUE" json:"email"`
	Password  string         `json:"password"`
	Role      string         `gorm:"type:enum('Admin','Customer');default:Customer" json:"role"`
	Status    string         `gorm:"type:enum('Pending','Active','Suspended');default:Active" json:"status"`
	OTP       string         `json:"otp"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type UserLogin struct {
	ID     uint   `json:"id" gorm:"primarykey"`
	Name   string `json:"name"`
	Email  string `gorm:"UNIQUE" json:"email"`
	Role   string `gorm:"type:enum('Admin','Customer');default:Customer" json:"role"`
	Status string `gorm:"type:enum('Pending','Active','Suspended');default:Active" json:"status"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}

type UserResponseSingle struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}

type UserResponseLogin struct {
	Status  bool      `json:"status"`
	Message string    `json:"message"`
	Data    UserLogin `json:"data"`
}
