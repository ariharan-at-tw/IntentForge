package repository

import "learn-commerce-service/internal/models"

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
