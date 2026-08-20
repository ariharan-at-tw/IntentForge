package repository

import (
	"fmt"
	"learn-commerce-service/internal/models"
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
