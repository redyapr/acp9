package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	UserID    int            `json:"userId"`
	Status    string         `json:"status" gorm:"type:enum('Paid','Unpaid','Canceled');default:Unpaid"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type TransactionDetail struct {
	ID            uint           `json:"id" gorm:"primarykey"`
	TransactionID int            `json:"transactionId"`
	ProductID     int            `json:"productId"`
	Qty           int            `json:"qty"`
	Price         int            `json:"price"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
