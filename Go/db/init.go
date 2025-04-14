package db

import (
	"github.com/vkazakevich/ebiznes/Go/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Cart{})

	return db
}
