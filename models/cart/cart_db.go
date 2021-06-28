package cart

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	ID     uint `gorm:"primarykey" json:"id"`
	UserId int  `json:"userId"`
}
