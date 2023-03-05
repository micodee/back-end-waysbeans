package controllers

import (
	"net/http"
	"strconv"
	"waysbeans/dto"
	"waysbeans/dto/result"
	"waysbeans/models"
	"waysbeans/repositories"

	"github.com/labstack/echo/v4"
)

// SETUP CONTROL STRUCT
type cartControl struct {
	CartRepository repositories.CartRepository
}

// SETUP CONTROL FUNCTION
func ControlCart(CartRepository repositories.CartRepository) *cartControl {
	return &cartControl{CartRepository}
}

// FUNCTION FIND CARTS
func (h *cartControl) FindCarts(c echo.Context) error {
	carts, err := h.CartRepository.FindCart()
	if err != nil {
		return c.JSON(http.StatusBadRequest, result.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, result.SuccessResult{Status: http.StatusOK, Data: carts})
}

// FUNCTION GET CART BY ID
func (h *cartControl) GetCart(c echo.Context) error {
	// get url param ID
	id, _ := strconv.Atoi(c.Param("id"))

	// repository get profile
	var cart models.Cart
	cart, err := h.CartRepository.GetCart(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, result.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, result.SuccessResult{Status: http.StatusOK, Data: convCart(cart)})
}

func convCart(u models.Cart) dto.CartResponse {
	return dto.CartResponse{
		OrderQty:  u.OrderQty,
		UserID:    u.UserID,
	}
}
