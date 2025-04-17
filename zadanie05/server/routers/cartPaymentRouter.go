package routers

import (
	"github.com/labstack/echo/v4"
	"productsServer/controllers"
)

func RegisterPaymentCartRoutes(e *echo.Echo) {
	e.GET("/payment", controllers.GetPaymentsWithCarts)
	e.GET("/payment/:id", controllers.GetPaymentWithCart)
	e.POST("/payment", controllers.CreatePaymentWithCart)
}
