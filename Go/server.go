package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/vkazakevich/ebiznes/Go/controllers"
	"github.com/vkazakevich/ebiznes/Go/db"
)

func main() {
	db := db.InitDatabase()

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())

	с := &controllers.Controller{DB: db}

	product := e.Group("/products")

	product.GET("", с.GetAllProduct)
	product.POST("", с.CreateProduct)
	product.GET("/:id", с.FindProduct)
	product.PUT("/:id", с.UpdateProduct)
	product.DELETE("/:id", с.DeleteProduct)

	e.Logger.Fatal(e.Start(":8000"))
}
