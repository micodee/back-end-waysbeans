package models

import "time"

type Product struct {
	ID          int            `json:"id"`
	Name        string         `json:"name" gorm:"type: varchar(255)"`
	Description string         `json:"description" gorm:"type: varchar(255)"`
	Price       int            `json:"price" gorm:"type: int"`
	Stock       int            `json:"stock" gorm:"type: int"`
	Photo       string         `json:"photo" gorm:"type: varchar(255)"`
	UserID      int            `json:"user_id" form:"user_id"`
	User        UsersToProduct `json:"user"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type ProductToCart struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Photo       string `json:"photo"`
	Stock       int    `json:"stock"`
}

func (ProductToCart) TableName() string {
	return "products"
}
