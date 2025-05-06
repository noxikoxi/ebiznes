package routers

import (
	"github.com/labstack/echo/v4"
	"zadanie08/controllers"
)

func LoginRoutes(e *echo.Echo) {
	e.POST("/login", controllers.LoginUser)

}
