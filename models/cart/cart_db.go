package cart

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserId int `json:"userId"`
}
