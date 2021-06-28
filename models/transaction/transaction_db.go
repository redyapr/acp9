package transaction

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	ID     uint   `gorm:"primarykey;autoIncrement" json:"id"`
	UserID int    `json:"userId"`
	Status string `json:"status" gorm:"type:enum('Paid','Unpaid','Canceled');default:Unpaid"`
}
