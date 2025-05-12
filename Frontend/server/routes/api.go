package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/vkazakevich/ebiznes/Go/controllers"
)

func ApiRoutes(e *echo.Echo, c *controllers.Controller) {
	product := e.Group("/products")
	product.GET("", c.GetAllProduct)
	product.POST("", c.CreateProduct)
	product.GET("/:id", c.FindProduct)
	product.PUT("/:id", c.UpdateProduct)
	product.DELETE("/:id", c.DeleteProduct)

	p := e.Group("/payments")
	p.POST("", c.CreatePayment)
}
