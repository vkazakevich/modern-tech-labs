package models

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCartModel(t *testing.T) {
	db := setupTestDB()

	t.Run("It creates Cart", func(t *testing.T) {
		c := &Cart{Quantity: 10, ProductID: 1, CustomerID: 1}
		result := db.Create(&c)

		assert.NoError(t, result.Error)
		assert.Equal(t, int64(1), result.RowsAffected)
		
		assert.NotNil(t, c.ID)
		assert.Equal(t, uint(10), c.Quantity)
	})
}

