package models

import "gorm.io/gorm"

type Category struct {
	ID       uint      `gorm:"primarykey"`
	Name     string    `json:"name"`
	Products []Product `json:"products,omitempty"`
}

func WithProducts(db *gorm.DB) *gorm.DB {
	return db.Preload("Products")
}

func HasProducts(db *gorm.DB) *gorm.DB {
	return db.Joins("JOIN products ON products.category_id = categories.id").
		Group("categories.id").
		Having("COUNT(products.id) > ?", 0)
}
