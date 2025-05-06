package routers

import (
	"github.com/labstack/echo/v4"
	"zadanie08/controllers"
)

func GithubRoutes(e *echo.Echo) {
	e.GET("/github/login", controllers.HandleGithubLogin)
	e.GET("/github/callback", controllers.HandleGithubCallback)
}
