package transaction

import "gorm.io/gorm"

type TransactionDetail struct {
	gorm.Model
	ID            uint `gorm:"primarykey;autoIncrement" json:"id"`
	TransactionID int  `json:"transactionId"`
	ProductID     int  `json:"productId"`
	DetailQTY     int  `json:"qty"`
	DetailPrice   int  `json:"price"`
}
