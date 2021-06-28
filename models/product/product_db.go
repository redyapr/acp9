package product

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	CategoryID int    `json:"categoryId"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Stockint   int    `json:"stock"`
}
