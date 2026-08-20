package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
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

func (h *ProductHandler) SearchProductsByName(
	w http.ResponseWriter,
	r *http.Request,
) {
	name := r.URL.Query().Get("name")

	if name == "" {
		http.Error(
			w,
			"name query parameter is required",
			http.StatusBadRequest,
		)
		return
	}

	products := h.productService.SearchProductsByName(name)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(
			w,
			"failed to encode products",
			http.StatusInternalServerError,
		)
	}
}

func (h *ProductHandler) GetProductsByCategory(
	w http.ResponseWriter,
	r *http.Request,
) {
	category := r.URL.Query().Get("category")

	if category == "" {
		http.Error(
			w,
			"category query parameter is required",
			http.StatusBadRequest,
		)
		return
	}

	products := h.productService.GetProductsByCategory(category)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(
			w,
			"failed to encode products",
			http.StatusInternalServerError,
		)
	}
}

func (h *ProductHandler) GetProductsByPriceRange(
	w http.ResponseWriter,
	r *http.Request,
) {
	query := r.URL.Query()

	var minPrice *float64
	var maxPrice *float64

	if value := query.Get("min_price"); value != "" {
		parsed, err := strconv.ParseFloat(value, 64)
		if err != nil {
			http.Error(
				w,
				"invalid min_price",
				http.StatusBadRequest,
			)
			return
		}

		minPrice = &parsed
	}

	if value := query.Get("max_price"); value != "" {
		parsed, err := strconv.ParseFloat(value, 64)
		if err != nil {
			http.Error(
				w,
				"invalid max_price",
				http.StatusBadRequest,
			)
			return
		}

		maxPrice = &parsed
	}

	if minPrice != nil && maxPrice != nil && *minPrice > *maxPrice {
		http.Error(
			w,
			"min_price cannot be greater than max_price",
			http.StatusBadRequest,
		)
		return
	}

	products := h.productService.GetProductsByPriceRange(
		minPrice,
		maxPrice,
	)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(
			w,
			"failed to encode products",
			http.StatusInternalServerError,
		)
	}
}
