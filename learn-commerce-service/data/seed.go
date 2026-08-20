package data

import "learn-commerce-service/internal/models"

func SeedProducts() []models.Product {
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
			Name:          "Logitech MX Keys",
			Category:      "accessories",
			Price:         10999,
			StockQuantity: 20,
		},
		{
			ID:            "P004",
			Name:          "Sony WH-1000XM6",
			Category:      "headphones",
			Price:         34999,
			StockQuantity: 0,
		},
	}
}
