package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

const DEFAULT_PORT = "1323"

func main() {

	// connect to the database
	var err error

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
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

		DB.Last(&createdProduct)

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "product created",
			"data":    createdProduct,
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

	var port string = os.Getenv("PORT")

	if port == "" {
		port = DEFAULT_PORT
	}

	var appPort string = fmt.Sprintf(":%s", port)

	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${method} ${path} ${status}",
	}))

	app.Logger.Fatal(app.Start(appPort))
}
