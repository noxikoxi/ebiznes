package routers

import (
	"github.com/labstack/echo/v4"
	"zadanie08/controllers"
)

func GoogleRoutes(e *echo.Echo) {
	e.GET("/google/login", controllers.HandleGoogleLogin)
	e.GET("/google/callback", controllers.HandleGoogleCallback)

}
