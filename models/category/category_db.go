package category

import (
	"gorm.io/gorm"
	"time"
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
