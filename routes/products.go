package routes

import (
	"waysbeans/controllers"
	"waysbeans/pkg/middleware"
	"waysbeans/pkg/mysql"
	"waysbeans/repositories"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Group) {
	productRepository := repositories.RepositoryProduct(mysql.ConnDB)
	h := controllers.ControlProduct(productRepository)

	e.GET("/products", h.FindProducts)
	e.GET("/product/:id", h.GetProducts)
	e.POST("/product", middleware.UploadFile(h.CreateProduct))
	e.PATCH("/product/:id", h.UpdateProduct)
	e.DELETE("/product/:id", h.DeleteProduct)
}