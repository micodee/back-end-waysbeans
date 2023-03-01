package repositories

import (
	"backEnd/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindProduct() ([]models.Product, error)
	GetProducts(ID int) (models.Product, error)
}

func RepositoryProduct(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindProduct() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *repository) GetProducts(ID int) (models.Product, error) {
	var products models.Product
	err := r.db.First(&products, ID).Error

	return products, err
}