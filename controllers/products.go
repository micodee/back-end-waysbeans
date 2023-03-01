package controllers

import (
	"net/http"
	"strconv"
	"waysbeans/dto"
	"waysbeans/dto/result"
	"waysbeans/models"
	"waysbeans/repositories"

	"github.com/go-playground/validator"
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

func (h *productControl) CreateProduct(c echo.Context) error {
	request := new(dto.CreateProductRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, result.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, result.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	// data form pattern submit to pattern entity db product
	product := models.Product{
		Name:        request.Name,
		Price:       request.Price,
		Description: request.Description,
		Stock:       request.Stock,
		Photo:       request.Photo,
	}

	data, err := h.ProductRepository.CreateProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, result.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, result.SuccessResult{Status: http.StatusOK, Data: convProduct(data)})
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