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

	h := &controllers.Controller{DB: db}

	product := e.Group("/products")

	product.GET("", h.GetAllProduct)
	product.POST("", h.CreateProduct)
	product.GET("/:id", h.FindProduct)
	product.PUT("/:id", h.UpdateProduct)
	product.DELETE("/:id", h.DeleteProduct)

	e.Logger.Fatal(e.Start(":8000"))
}
