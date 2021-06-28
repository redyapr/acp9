package category

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID          uint   `gorm:"primarykey;autoIncrement" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
