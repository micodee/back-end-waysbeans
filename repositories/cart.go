package repositories

import (
	"waysbeans/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	FindCart() ([]models.Cart, error)
}

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db}
}



func (r *repository) FindCart() ([]models.Cart, error) {
	var cart []models.Cart
	err := r.db.Preload("User").Preload("Product").Find(&cart).Error
	return cart, err
}