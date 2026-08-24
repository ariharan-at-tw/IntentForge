package repository

import (
	"testing"

	"learn-commerce-service/internal/models"
)

func TestProductRepository_GetByID(t *testing.T) {
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

	repository := NewProductRepository(products)

	product, err := repository.GetByID("P001")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if product.ID != "P001" {
		t.Errorf("expected product ID P001, got %s", product.ID)
	}
}

func TestProductRepository_GetByID_NotFound(t *testing.T) {
	products := []models.Product{
		{
			ID:            "P001",
			Name:          "iPhone 17",
			Category:      "smartphone",
			Price:         79999,
			StockQuantity: 12,
		},
	}

	repository := NewProductRepository(products)

	product, err := repository.GetByID("P999")

	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	if product != nil {
		t.Errorf("expected product to be nil, got %+v", product)
	}
}

func testProducts() []models.Product {
	return []models.Product{
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
		{
			ID:            "P003",
			Name:          "Samsung Galaxy",
			Category:      "smartphone",
			Price:         49999,
			StockQuantity: 8,
		},
	}
}

func float64Ptr(value float64) *float64 {
	return &value
}

func TestProductRepository_GetProducts(t *testing.T) {
	tests := []struct {
		name          string
		filter        models.ProductFilter
		expectedCount int
	}{
		{
			name:          "returns all products with empty filter",
			filter:        models.ProductFilter{},
			expectedCount: 3,
		},
		{
			name: "filters by name",
			filter: models.ProductFilter{
				Name: "iphone",
			},
			expectedCount: 1,
		},
		{
			name: "filters by category",
			filter: models.ProductFilter{
				Category: "smartphone",
			},
			expectedCount: 2,
		},
		{
			name: "filters by minimum price",
			filter: models.ProductFilter{
				MinPrice: float64Ptr(50000),
			},
			expectedCount: 2,
		},
		{
			name: "filters by maximum price",
			filter: models.ProductFilter{
				MaxPrice: float64Ptr(80000),
			},
			expectedCount: 2,
		},
		{
			name: "filters by price range",
			filter: models.ProductFilter{
				MinPrice: float64Ptr(50000),
				MaxPrice: float64Ptr(100000),
			},
			expectedCount: 2,
		},
		{
			name: "combines category and maximum price",
			filter: models.ProductFilter{
				Category: "smartphone",
				MaxPrice: float64Ptr(60000),
			},
			expectedCount: 1,
		},
		{
			name: "returns empty result when no products match",
			filter: models.ProductFilter{
				Category: "tablet",
			},
			expectedCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := NewProductRepository(testProducts())

			products := repository.GetProducts(tt.filter)

			if len(products) != tt.expectedCount {
				t.Errorf(
					"expected %d products, got %d",
					tt.expectedCount,
					len(products),
				)
			}
		})
	}
}
