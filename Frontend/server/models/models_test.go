package models

import (
	"gorm.io/gorm"
	"github.com/vkazakevich/ebiznes/Go/db"
)

func setupTestDB() *gorm.DB {
	db := db.InitTestDB()
	
	err := db.AutoMigrate(
		&Product{},
		&Cart{},
		&Category{},
		&Customer{},
		&Payment{},
	)

	if err != nil {
		panic("failed to migrate database")
	}

	return db
}