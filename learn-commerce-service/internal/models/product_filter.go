package models

type ProductFilter struct {
	Name     string
	Category string
	MinPrice *float64
	MaxPrice *float64
}
