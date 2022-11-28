package main

import (
	"simple-api/database"
	"simple-api/routes"

	"github.com/labstack/echo/v4"
)

//	@title			Simple Products API
//	@version		1.0
//	@description	This is a simple Products API
func main() {
	database.InitDB()

	app := echo.New()

	routes.SetupRoutes(app)

	app.Logger.Fatal(app.Start(":1323"))
}
