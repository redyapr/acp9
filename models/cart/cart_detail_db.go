package cart

import "gorm.io/gorm"

type CartDetail struct {
	gorm.Model
	ID        uint `gorm:"primarykey" json:"id"`
	CartId    int  `json:"cartId"`
	ProductID int  `json:"productId"`
	Qty       int  `json:"qty"`
}
