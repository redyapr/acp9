package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Slug        string         `json:"slug"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type CategoryResponse struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Data    []Category `json:"data"`
}
