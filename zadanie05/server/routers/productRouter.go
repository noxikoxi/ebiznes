package routers

import (
	"github.com/labstack/echo/v4"
	"productsServer/controllers"
)

func RegisterProductRoutes(e *echo.Echo) {
	e.GET("/products", controllers.GetProducts)
	e.GET("/products/:id", controllers.GetProduct)
	e.POST("/products", controllers.CreateProduct)
	e.DELETE("/products/:id", controllers.DeleteProduct)
}
