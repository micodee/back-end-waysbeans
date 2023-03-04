package models

type Cart struct {
	ID     int `json:"id" gorm:"primary_key:auto_increment"`
	Qty    int `json:"order_qty" gorm:"type: int"`
	Amount int `json:"amount" gorm:"type: int"`
}
