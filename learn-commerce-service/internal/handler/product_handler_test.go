package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"encoding/json"

	"github.com/gin-gonic/gin"

	"learn-commerce-service/internal/models"
	"learn-commerce-service/internal/repository"
	"learn-commerce-service/internal/service"
)

func setupProductHandler() *ProductHandler {
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
	productService := service.NewProductService(productRepository)

	return NewProductHandler(productService)
}

func TestProductHandler_GetProducts(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()

	handler := setupProductHandler()

	router.GET("/products", handler.GetProducts)

	request := httptest.NewRequest(
		http.MethodGet,
		"/products",
		nil,
	)

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusOK,
			response.Code,
		)
	}
}

func TestProductHandler_GetProductByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()

	handler := setupProductHandler()

	router.GET("/products/:id", handler.GetProductByID)

	request := httptest.NewRequest(
		http.MethodGet,
		"/products/P001",
		nil,
	)

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusOK,
			response.Code,
		)
	}

	var product models.Product

	if err := json.Unmarshal(
		response.Body.Bytes(),
		&product,
	); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if product.ID != "P001" {
		t.Errorf(
			"expected product P001, got %s",
			product.ID,
		)
	}
}

func TestProductHandler_GetProductByID_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()

	handler := setupProductHandler()

	router.GET("/products/:id", handler.GetProductByID)

	request := httptest.NewRequest(
		http.MethodGet,
		"/products/P999",
		nil,
	)

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusNotFound {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusNotFound,
			response.Code,
		)
	}
}
