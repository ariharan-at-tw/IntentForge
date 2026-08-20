package service

import (
	"learn-commerce-service/internal/models"
	"learn-commerce-service/internal/repository"
)

type ProductService struct {
	productRepository *repository.ProductRepository
}

func NewProductService(
	productRepository *repository.ProductRepository,
) *ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (s *ProductService) GetAllProducts() []models.Product {
	return s.productRepository.GetAll()
}

func (s *ProductService) GetProductByID(id string) (*models.Product, error) {
	return s.productRepository.GetByID(id)
}

func (s *ProductService) SearchProductsByName(name string) []models.Product {
	return s.productRepository.SearchByName(name)
}

func (s *ProductService) GetProductsByCategory(category string) []models.Product {
	return s.productRepository.GetByCategory(category)
}

func (s *ProductService) GetProductsByPriceRange(
	minPrice *float64,
	maxPrice *float64,
) []models.Product {
	return s.productRepository.GetByPriceRange(minPrice, maxPrice)
}
