package main

import (
	"learn-commerce-service/data"
	"learn-commerce-service/internal/handler"
	"learn-commerce-service/internal/repository"
	"learn-commerce-service/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	products := data.SeedProducts()

	productRepository := repository.NewProductRepository(products)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	router := gin.Default()

	router.GET("/products", productHandler.GetProducts)
	router.GET("/products/:id", productHandler.GetProductByID)

	router.Run(":8080")
}
