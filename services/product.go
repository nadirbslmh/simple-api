package services

import (
	"simple-api/database"
	"simple-api/models"
)

func CreateProduct(product models.Product) (models.Product, error) {
	if err := database.DB.Create(&product).Error; err != nil {
		return models.Product{}, err
	}

	var createdProduct models.Product

	database.DB.Last(&createdProduct)

	return createdProduct, nil
}

func GetProducts() []models.Product {
	var products []models.Product

	if err := database.DB.Find(&products).Error; err != nil {
		return []models.Product{}
	}

	return products
}

func GetProductsByName(keyword string) []models.Product {
	var products []models.Product

	productName := "%" + keyword + "%"

	if err := database.DB.Where("name LIKE ?", productName).Find(&products).Error; err != nil {
		return []models.Product{}
	}

	return products
}
