package transaction

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	ID                uint   `gorm:"primarykey;autoIncrement" json:"id"`
	UserID            int    `json:"userId"`
	TransactionStatus string `json:"status"`
}
