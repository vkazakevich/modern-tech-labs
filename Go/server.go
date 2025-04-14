package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/vkazakevich/ebiznes/Go/controllers"
	"github.com/vkazakevich/ebiznes/Go/db"
	"github.com/vkazakevich/ebiznes/Go/routes"
)

func main() {
	db := db.InitDatabase()

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	с := &controllers.Controller{DB: db}
	routes.ApiRoutes(e, с)

	e.Logger.Fatal(e.Start(":8000"))
}
