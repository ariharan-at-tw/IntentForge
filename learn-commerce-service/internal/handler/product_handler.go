package handler

import (
	"encoding/json"
	"net/http"

	"learn-commerce-service/internal/service"
)

type ProductHandler struct {
	productService *service.ProductService
}

func NewProductHandler(
	productService *service.ProductService,
) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (h *ProductHandler) GetProducts(
	w http.ResponseWriter,
	r *http.Request,
) {
	products := h.productService.GetAllProducts()

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(
			w,
			"failed to encode products",
			http.StatusInternalServerError,
		)
	}
}
