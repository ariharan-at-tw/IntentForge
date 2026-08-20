package handler

import (
	"encoding/json"
	"net/http"
	"strings"

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

func (h *ProductHandler) GetProductByID(
	w http.ResponseWriter,
	r *http.Request,
) {
	id := strings.TrimPrefix(r.URL.Path, "/products/")

	product, err := h.productService.GetProductByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(product); err != nil {
		http.Error(
			w,
			"failed to encode product",
			http.StatusInternalServerError,
		)
	}
}
