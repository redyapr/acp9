package models

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	UserId    int            `json:"user_id"`
	ProductID int            `json:"product_id"`
	Qty       int            `json:"qty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	// Product   Product        `gorm:"foreignKey:ID;reference:product_id" json:"product"`
}

type CartResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    []Cart `json:"data"`
}

type CartResponseSingle struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    Cart   `json:"data"`
}
