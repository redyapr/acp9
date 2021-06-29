package transaction

import "gorm.io/gorm"

type TransactionDetail struct {
	gorm.Model
	TransactionID int `json:"transactionId"`
	ProductID     int `json:"productId"`
	Qty           int `json:"qty"`
	Price         int `json:"price"`
}
