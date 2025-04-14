package utils

import (
	"github.com/vkazakevich/ebiznes/Go/models"
	"gorm.io/gorm"
)

func FillDBSeeds(db *gorm.DB) {
	// Categories
	db.Create(&models.Category{Name: "Laptops"})
	db.Create(&models.Category{Name: "Mobile Phones"})
	db.Create(&models.Category{Name: "Tablets"})

	// Products
	var c models.Category
	db.First(&c, "name = ?", "Laptops")
	db.Create(&models.Product{Name: "Macbook", Price: 1099, Quantity: 10, Category: c})
	db.Create(&models.Product{Name: "Macbook Air", Price: 899, Quantity: 2, Category: c})
	db.Create(&models.Product{Name: "Macbook PRO", Price: 1299, Quantity: 5, Category: c})
}
