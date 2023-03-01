package controllers

import (
	"backEnd/dto/result"
	"backEnd/repositories"
	"net/http"

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