package category

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
}
