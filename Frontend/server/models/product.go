package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name       string   `json:"name"`
	Price      uint     `json:"price"`
	Quantity   uint     `json:"quantity"`
	CategoryID uint     `json:"category_id"`
	Category   Category `json:"-"`
}
