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

func TestProductHandler_GetProducts_WithFilters(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name          string
		query         string
		expectedCount int
	}{
		{
			name:          "filter by name",
			query:         "?name=iphone",
			expectedCount: 1,
		},
		{
			name:          "filter by category",
			query:         "?category=smartphone",
			expectedCount: 1,
		},
		{
			name:          "filter by minimum price",
			query:         "?min_price=80000",
			expectedCount: 1,
		},
		{
			name:          "filter by maximum price",
			query:         "?max_price=80000",
			expectedCount: 1,
		},
		{
			name:          "filter by price range",
			query:         "?min_price=70000&max_price=90000",
			expectedCount: 1,
		},
		{
			name:          "combine category and price",
			query:         "?category=smartphone&max_price=80000",
			expectedCount: 1,
		},
		{
			name:          "no matching products",
			query:         "?category=tablet",
			expectedCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.New()
			handler := setupProductHandler()

			router.GET("/products", handler.GetProducts)

			request := httptest.NewRequest(
				http.MethodGet,
				"/products"+tt.query,
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

			var products []models.Product

			if err := json.Unmarshal(
				response.Body.Bytes(),
				&products,
			); err != nil {
				t.Fatalf("failed to decode response: %v", err)
			}

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

func TestProductHandler_GetProducts_InvalidPrice(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	handler := setupProductHandler()

	router.GET("/products", handler.GetProducts)

	request := httptest.NewRequest(
		http.MethodGet,
		"/products?min_price=abc",
		nil,
	)

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusBadRequest,
			response.Code,
		)
	}
}

func TestProductHandler_GetProducts_InvalidPriceRange(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	handler := setupProductHandler()

	router.GET("/products", handler.GetProducts)

	request := httptest.NewRequest(
		http.MethodGet,
		"/products?min_price=90000&max_price=50000",
		nil,
	)

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusBadRequest,
			response.Code,
		)
	}
}
