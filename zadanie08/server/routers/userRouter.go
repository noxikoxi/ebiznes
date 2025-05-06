package routers

import (
	"github.com/labstack/echo/v4"
	"zadanie08/controllers"
)

func UserRoutes(e *echo.Echo) {
	e.GET("/user", controllers.GetUserData)
}
