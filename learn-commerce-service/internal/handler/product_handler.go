package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"learn-commerce-service/internal/models"
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

func (h *ProductHandler) GetProducts(c *gin.Context) {
	filter := models.ProductFilter{
		Name:     c.Query("name"),
		Category: c.Query("category"),
	}

	if value := c.Query("min_price"); value != "" {
		minPrice, err := strconv.ParseFloat(value, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid min_price",
			})
			return
		}

		filter.MinPrice = &minPrice
	}

	if value := c.Query("max_price"); value != "" {
		maxPrice, err := strconv.ParseFloat(value, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid max_price",
			})
			return
		}

		filter.MaxPrice = &maxPrice
	}

	if filter.MinPrice != nil &&
		filter.MaxPrice != nil &&
		*filter.MinPrice > *filter.MaxPrice {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "min_price cannot be greater than max_price",
		})
		return
	}

	products := h.productService.GetProducts(filter)

	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id := c.Param("id")

	product, err := h.productService.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)
}
