package service

import (
	"testing"

	"learn-commerce-service/internal/models"
	"learn-commerce-service/internal/repository"
)

func TestProductService_GetProductByID(t *testing.T) {
	products := []models.Product{
		{
			ID:            "P001",
			Name:          "iPhone 17",
			Category:      "smartphone",
			Price:         79999,
			StockQuantity: 12,
		},
	}

	productRepository := repository.NewProductRepository(products)
	productService := NewProductService(productRepository)

	product, err := productService.GetProductByID("P001")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if product == nil {
		t.Fatal("expected product, got nil")
	}

	if product.ID != "P001" {
		t.Errorf("expected product ID P001, got %s", product.ID)
	}
}

func TestProductService_GetProductByID_NotFound(t *testing.T) {
	productRepository := repository.NewProductRepository(nil)
	productService := NewProductService(productRepository)

	product, err := productService.GetProductByID("P999")

	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	if product != nil {
		t.Errorf("expected product to be nil, got %+v", product)
	}
}

func TestProductService_GetProducts(t *testing.T) {
	products := []models.Product{
		{
			ID:            "P001",
			Name:          "iPhone 17",
			Category:      "smartphone",
			Price:         79999,
			StockQuantity: 12,
		},
		{
			ID:            "P002",
			Name:          "MacBook Air",
			Category:      "laptop",
			Price:         99999,
			StockQuantity: 5,
		},
	}

	productRepository := repository.NewProductRepository(products)
	productService := NewProductService(productRepository)

	filter := models.ProductFilter{
		Category: "laptop",
	}

	result := productService.GetProducts(filter)

	if len(result) != 1 {
		t.Fatalf("expected 1 product, got %d", len(result))
	}

	if result[0].ID != "P002" {
		t.Errorf("expected P002, got %s", result[0].ID)
	}
}
