package cart

import "gorm.io/gorm"

type CartDetail struct {
	gorm.Model
	CartId    int `json:"cartId"`
	ProductID int `json:"productId"`
	Qty       int `json:"qty"`
}
