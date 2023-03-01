package controllers

import (
	"backEnd/dto"
	"backEnd/dto/result"
	"backEnd/models"
	"backEnd/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type productControl struct {
	ProductRepository repositories.ProductRepository
}

func ControlProduct(ProductRepository repositories.ProductRepository) *productControl {
	return &productControl{ProductRepository}
}



func (h *productControl) FindProducts(c echo.Context) error {
	products, err := h.ProductRepository.FindProduct()
	if err != nil {
		return c.JSON(http.StatusBadRequest, result.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, result.SuccessResult{Status: http.StatusOK, Data: products})
}

func (h *productControl) GetProducts(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	products, err := h.ProductRepository.GetProducts(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, result.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, result.SuccessResult{Status: http.StatusOK, Data: convProduct(products)})
}

func convProduct(u models.Product) dto.ProductResponse {
	return dto.ProductResponse{
		ID:          u.ID,
		Name:        u.Name,
		Price:       u.Price,
		Description: u.Description,
		Stock:       u.Stock,
		Photo:       u.Photo,
	}
}