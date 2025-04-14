package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Quantity uint `json:"quantity"`

	ProductID uint
	Product   Product

	CustomerID uint
	Customer   Product
}

func WithProduct(db *gorm.DB) *gorm.DB {
	return db.Preload("Product")
}
