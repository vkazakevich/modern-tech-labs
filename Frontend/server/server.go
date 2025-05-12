package main

import (
	"flag"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/vkazakevich/ebiznes/Go/controllers"
	"github.com/vkazakevich/ebiznes/Go/db"
	"github.com/vkazakevich/ebiznes/Go/routes"
	"github.com/vkazakevich/ebiznes/Go/utils"
)

func main() {
	var fillSeeds bool
	flag.BoolVar(&fillSeeds, "seed", false, "a bool")
	flag.Parse()

	db := db.InitDB()
	utils.MigrateDB(db)

	if fillSeeds {
		utils.FillDBSeeds(db)
	}

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"http://localhost:5173"},
        AllowMethods: []string{"GET", "POST"},
    }))

	c := &controllers.Controller{DB: db}
	routes.ApiRoutes(e, c)

	e.Logger.Fatal(e.Start(":8000"))
}
