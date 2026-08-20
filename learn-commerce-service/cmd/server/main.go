package main

import (
	"fmt"
	"net/http"

	"learn-commerce-service/data"
	"learn-commerce-service/internal/handler"
	"learn-commerce-service/internal/repository"
	"learn-commerce-service/internal/service"
)

func main() {
	products := data.SeedProducts()

	productRepository :=
		repository.NewProductRepository(products)

	productService :=
		service.NewProductService(productRepository)

	productHandler :=
		handler.NewProductHandler(productService)

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("name") != "" {
			productHandler.SearchProductsByName(w, r)
			return
		}

		if r.URL.Query().Get("category") != "" {
			productHandler.GetProductsByCategory(w, r)
			return
		}

		productHandler.GetProducts(w, r)
	})
	http.HandleFunc("/products/", productHandler.GetProductByID)

	fmt.Println(
		"Commerce service running on http://localhost:8080",
	)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
