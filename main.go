package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       int            `json:"price"`
	Quantity    int            `json:"quantity"`
}

var DB *gorm.DB

func main() {

	// connect to the database
	var err error

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		"root",
		"",
		"localhost",
		"3306",
		"echo_sample",
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("error: ", err)
	}

	DB.AutoMigrate(&Product{})

	log.Println("connected to the database")

	app := echo.New()

	app.POST("/products", func(c echo.Context) error {
		var product Product

		if err := c.Bind(&product); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "bad request")
		}

		if err := DB.Create(&product).Error; err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "cannot insert data")
		}

		var createdProduct Product

		created := DB.Last(&createdProduct)

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "product created",
			"data":    created,
		})

	})

	app.GET("/products", func(c echo.Context) error {
		var products []Product

		if err := DB.Find(&products).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "bad request")
		}

		return c.JSON(http.StatusOK, map[string]any{
			"message": "all products",
			"data":    products,
		})
	})

	app.Logger.Fatal(app.Start(":1323"))
}
