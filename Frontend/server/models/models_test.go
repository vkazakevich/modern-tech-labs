package models

import (
	"gorm.io/gorm"
	"github.com/vkazakevich/ebiznes/Go/db"
)

func setupTestDB() *gorm.DB {
	db := db.InitTestDB()
	
	db.AutoMigrate(
		&Product{},
		&Cart{},
		&Category{},
		&Customer{},
		&Payment{},
	)

	return db
}