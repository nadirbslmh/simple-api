package controllers

import (
	"net/http"
	"simple-api/models"
	"simple-api/services"

	"github.com/labstack/echo/v4"
)

type Response[T any] struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// @Summary		create a product
// @Description	create a new product data
// @Accept			json
// @Produce		json
// @Param			product	body		models.Product					true	"Create product"
// @Success		201		{object}	Response{data=models.Product}	"product created"
// @Router			/products [post]
func CreateProduct(c echo.Context) error {
	var product models.Product

	if err := c.Bind(&product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}

	createdProduct, err := services.CreateProduct(product)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create a product")
	}

	return c.JSON(http.StatusCreated, Response[models.Product]{
		Message: "product created",
		Data:    createdProduct,
	})
}

// @Summary		get all products
// @Description	get all products data
// @Produce		json
// @Success		200	{object}	Response{data=[]models.Product}	"all products"
// @Router			/products [get]
func GetAll(c echo.Context) error {
	var products []models.Product

	productKeyword := c.QueryParam("keyword")

	if productKeyword != "" {
		products = services.GetProductsByName(productKeyword)
		return c.JSON(http.StatusOK, Response[[]models.Product]{
			Message: "all products",
			Data:    products,
		})
	}

	products = services.GetProducts()

	return c.JSON(http.StatusOK, Response[[]models.Product]{
		Message: "all products",
		Data:    products,
	})
}
