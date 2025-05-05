package utils

import (
	"testing"
	"gorm.io/gorm"

	"github.com/stretchr/testify/assert"
	"github.com/vkazakevich/ebiznes/Go/models"
	"github.com/vkazakevich/ebiznes/Go/db"
)

func setupTestDB() *gorm.DB {
	db := db.InitTestDB()
	MigrateDB(db)
	return db
}

func TestFillDBSeeds(t *testing.T) {
	db := setupTestDB()
	FillDBSeeds(db)

	var count int64
	db.Model(&models.Category{}).Count(&count)
	assert.Equal(t, int64(3), count)

	db.Model(&models.Product{}).Count(&count)
	assert.Equal(t, int64(3), count)
}