package product

import (
	"acp9-redy-gigih/models/category"
	"gorm.io/gorm"
	"time"
)

type ProductByCategory struct {
	ID         uint           `json:"id"`
	//CategoryID int            `json:"category_id"`
	Name       string         `json:"name"`
	Price      int            `json:"price"`
	Stockint   int            `json:"stock"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Category category.Category `gorm:"foreignKey:ID;reference:category_id" json:"category"`
}

type ProductResponse struct {
	Status bool `json:"status"`
	Message string `json:"message"`
	Data []Product `json:"data"`
}

type ProductResponseSingle struct {
	Status bool `json:"status"`
	Message string `json:"message"`
	Data Product `json:"data"`
}
