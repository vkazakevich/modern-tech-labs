package utils

import (
	"github.com/vkazakevich/ebiznes/Go/models"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Product{},
		&models.Cart{},
		&models.Category{},
		&models.Customer{},
		&models.Payment{},
	)

	if err != nil {
		panic("failed to migrate database")
	}
}