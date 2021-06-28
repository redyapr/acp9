package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primarykey" json:"id"`
	Name     string `json:"name"`
	Email    string `gorm:"UNIQUE" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Status   string `json:"status"`
	OTP      string `json:"otp"`
}
