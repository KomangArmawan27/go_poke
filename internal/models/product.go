package models

import "gorm.io/gorm"

// Product represents a product entity
type Product struct {
	gorm.Model
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
