package routers

import (
	"github.com/labstack/echo/v4"
	"zadanie08/controllers"
)

func RegisterRoutes(e *echo.Echo) {
	e.POST("/register", controllers.RegisterUser)
}
