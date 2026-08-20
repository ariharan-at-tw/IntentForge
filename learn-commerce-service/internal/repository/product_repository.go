package repository

import (
	"fmt"
	"learn-commerce-service/internal/models"
	"strings"
)

type ProductRepository struct {
	products []models.Product
}

func NewProductRepository(products []models.Product) *ProductRepository {
	return &ProductRepository{
		products: products,
	}
}

func (r *ProductRepository) GetAll() []models.Product {
	return r.products
}

func (r *ProductRepository) GetByID(id string) (*models.Product, error) {
	for _, product := range r.products {
		if product.ID == id {
			return &product, nil
		}
	}

	return nil, fmt.Errorf("product not found")
}

func (r *ProductRepository) SearchByName(name string) []models.Product {
	var results []models.Product

	for _, product := range r.products {
		if strings.Contains(
			strings.ToLower(product.Name),
			strings.ToLower(name),
		) {
			results = append(results, product)
		}
	}

	return results
}

func (r *ProductRepository) GetByCategory(category string) []models.Product {
	var results []models.Product

	for _, product := range r.products {
		if strings.EqualFold(product.Category, category) {
			results = append(results, product)
		}
	}

	return results
}
