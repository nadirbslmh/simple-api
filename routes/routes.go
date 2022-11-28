package routes

import (
	"simple-api/controllers"

	_ "simple-api/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func SetupRoutes(e *echo.Echo) {
	// swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/products", controllers.CreateProduct)
	e.GET("/products", controllers.GetAll)
}
