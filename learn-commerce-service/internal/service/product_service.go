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
