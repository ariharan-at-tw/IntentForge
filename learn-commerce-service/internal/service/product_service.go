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

func (s *ProductService) GetProducts(
	filter models.ProductFilter,
) []models.Product {
	return s.productRepository.GetProducts(filter)
}

func (s *ProductService) GetProductByID(id string) (*models.Product, error) {
	return s.productRepository.GetByID(id)
}
