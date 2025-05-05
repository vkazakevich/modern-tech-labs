package utils

import (
	"github.com/vkazakevich/ebiznes/Go/models"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(
		&models.Product{},
		&models.Cart{},
		&models.Category{},
		&models.Customer{},
		&models.Payment{},
	)
}