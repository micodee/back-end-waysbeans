package routes

import (
	"waysbeans/controllers"
	"waysbeans/pkg/mysql"
	"waysbeans/repositories"

	"github.com/labstack/echo/v4"
)

func ProfileRoutes(e *echo.Group) {
	profileRepository := repositories.RepositoryProfile(mysql.ConnDB)
	h := controllers.ControlProfile(profileRepository)

	e.GET("/profiles", h.FindProfile)
	e.GET("/profile/:id", h.GetProfile)
}
