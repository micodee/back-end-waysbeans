package controllers

import (
	"net/http"
	"waysbeans/dto/result"
	"waysbeans/repositories"

	"github.com/labstack/echo/v4"
)

// SETUP CONTROL STRUCT
type cartControl struct {
	CartRepository repositories.CartRepository
}

//	SETUP CONTROL FUNCTION
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