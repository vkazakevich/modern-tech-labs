package dto

import (
	"gorm.io/gorm"

	"github.com/vkazakevich/ebiznes/Go/db"
	"github.com/vkazakevich/ebiznes/Go/utils"
)

func setupTestDB() *gorm.DB {
	db := db.InitTestDB()
	utils.MigrateDB(db)

	return db
}