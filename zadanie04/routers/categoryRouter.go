package routers

import (
	"github.com/labstack/echo/v4"
	"zadanie04/controllers"
)

func RegisterCategoryRoutes(e *echo.Echo) {
	e.GET("/categories", controllers.GetCategories)
	e.GET("/categories/:id", controllers.GetCategory)
	e.POST("/categories", controllers.CreateCategory)
	e.PUT("/categories/:id", controllers.UpdateCategory)
	e.DELETE("/categories/:id", controllers.DeleteCategory)
}
