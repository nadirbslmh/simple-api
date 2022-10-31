package controllers

import (
	"net/http"
	"simple-api/models"
	"simple-api/services"

	"github.com/labstack/echo/v4"
)

func CreateProduct(c echo.Context) error {
	var product models.Product

	if err := c.Bind(&product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}

	createdProduct, err := services.CreateProduct(product)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create a product")
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"message": "product created",
		"data":    createdProduct,
	})
}

func GetAll(c echo.Context) error {
	products := services.GetProducts()

	return c.JSON(http.StatusOK, map[string]any{
		"message": "all products",
		"data":    products,
	})
}
