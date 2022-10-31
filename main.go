package main

import (
	"simple-api/database"
	"simple-api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDB()

	app := echo.New()

	routes.SetupRoutes(app)

	app.Logger.Fatal(app.Start(":1323"))
}
