package transaction

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UserID int    `json:"userId"`
	Status string `json:"status" gorm:"type:enum('Paid','Unpaid','Canceled');default:Unpaid"`
}
