package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/vkazakevich/ebiznes/Go/controllers"
)

func ApiRoutes(e *echo.Echo, с *controllers.Controller) {
	product := e.Group("/products")
	product.GET("", с.GetAllProduct)
	product.POST("", с.CreateProduct)
	product.GET("/:id", с.FindProduct)
	product.PUT("/:id", с.UpdateProduct)
	product.DELETE("/:id", с.DeleteProduct)

	cart := e.Group("/cart")
	cart.GET("", с.GetAllCartItems)
	cart.POST("", с.CreateCartItem)
	cart.GET("/:id", с.FindCartItem)
	cart.PUT("/:id", с.UpdateCartItem)
	cart.DELETE("/:id", с.DeleteCartItem)
}
