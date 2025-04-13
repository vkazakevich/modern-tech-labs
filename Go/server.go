package main

import (
	// "net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/vkazakevich/ebiznes/Go/controllers"
)

func main() {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())

	product := e.Group("/products")

	product.GET("", controllers.GetAllProduct)
	product.POST("", controllers.CreateProduct)
	product.GET("/:id", controllers.FindProduct)
	product.PUT("/:id", controllers.UpdateProduct)
	product.DELETE("/:id", controllers.DeleteProduct)

	e.Logger.Fatal(e.Start(":8000"))
}
