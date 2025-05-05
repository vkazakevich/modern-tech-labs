package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	if err != nil {
		panic("failed to connect to the test database")
	}
	
	return db
}
