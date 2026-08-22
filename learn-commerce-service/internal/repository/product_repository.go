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

func (r *ProductRepository) GetByID(id string) (*models.Product, error) {
	for _, product := range r.products {
		if product.ID == id {
			return &product, nil
		}
	}

	return nil, fmt.Errorf("product not found")
}

func (r *ProductRepository) GetProducts(
	filter models.ProductFilter,
) []models.Product {
	var results []models.Product

	for _, product := range r.products {
		if filter.Name != "" &&
			!strings.Contains(
				strings.ToLower(product.Name),
				strings.ToLower(filter.Name),
			) {
			continue
		}

		if filter.Category != "" &&
			!strings.EqualFold(
				product.Category,
				filter.Category,
			) {
			continue
		}

		if filter.MinPrice != nil &&
			product.Price < *filter.MinPrice {
			continue
		}

		if filter.MaxPrice != nil &&
			product.Price > *filter.MaxPrice {
			continue
		}

		results = append(results, product)
	}

	return results
}
