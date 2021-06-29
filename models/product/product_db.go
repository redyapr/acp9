package product

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID         uint           `json:"id" gorm:"primarykey"`
	CategoryID int            `json:"category_id"`
	Name       string         `json:"name"`
	Price      int            `json:"price"`
	Stockint   int            `json:"stock"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
