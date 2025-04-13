package routers

import (
	"github.com/labstack/echo/v4"
	"zadanie04/controllers"
)

func RegisterCartRoutes(e *echo.Echo) {
	e.GET("/carts", controllers.GetCarts)
	e.GET("/carts/:id", controllers.GetCart)
	e.POST("/carts", controllers.CreateCart)
	e.DELETE("/carts/:id", controllers.DeleteCart)

	// CartItems
	e.PUT("/carts/:cart_id/:product_id/:quantity", controllers.UpdateCartItem)
	e.POST("/carts/:cart_id/:product_id/:quantity", controllers.AddCartItem)
	e.DELETE("/carts/:cart_id/:product_id", controllers.DeleteCartItem)
}
