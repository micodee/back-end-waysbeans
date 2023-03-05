package models

type Cart struct {
	ID        int           `json:"id" gorm:"primary_key:auto_increment"`
	UserID    int           `json:"user_id"`
	User      UsersToCart `json:"user"`
	ProductID int           `json:"product_id" gorm:"type: int"`
	Product   ProductToCart `json:"product"`
	Qty       int           `json:"order_qty" gorm:"type: int"`
}

type CartUserResponse struct {
	ProductID int `json:"product_id" gorm:"type: int"`
	Qty       int `json:"order_qty" gorm:"type: int"`
	UserID    int `json:"-"`
}

type CartProductResponse struct {
	ProductID int `json:"-"`
	Qty       int `json:"order_qty" gorm:"type: int"`
	UserID    int `json:"user_id" gorm:"type: int"`
}

func (CartUserResponse) TableName() string {
	return "carts"
}

func (CartProductResponse) TableName() string {
	return "carts"
}
