package models

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCategoryModel(t *testing.T) {
	db := setupTestDB()

	t.Run("It creates Category", func(t *testing.T) {
		c := &Category{Name: "Test Category"}
		result := db.Create(&c)

		assert.NoError(t, result.Error)
		assert.Equal(t, int64(1), result.RowsAffected)
		
		assert.NotNil(t, c.ID)
		assert.Equal(t, "Test Category", c.Name)
	})
}

