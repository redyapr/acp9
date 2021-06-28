package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primarykey;autoIncrement" json:"id"`
	Name     string `json:"name"`
	Email    string `gorm:"UNIQUE" json:"email"`
	Password string `json:"password"`
	Role     string `gorm:"type:enum('Admin','Customer');default:Customer" json:"role"`
	Status   string `gorm:"type:enum('Pending','Active','Suspended');default:Pending" json:"status"`
	OTP      string `json:"otp"`
}
