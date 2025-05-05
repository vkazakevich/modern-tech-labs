package models

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestProductModel(t *testing.T) {
	db := setupTestDB()

	t.Run("It creates Product", func(t *testing.T) {
		p := &Product{Name: "Test Product"}
		result := db.Create(&p)

		assert.NoError(t, result.Error)
		assert.Equal(t, int64(1), result.RowsAffected)
		assert.NotNil(t, p.ID)
	})


	t.Run("It updates Product", func(t *testing.T) {
		p := &Product{Name: "Test Product"}
		db.Create(&p)

		p.Name = "Updated Test Product"
		p.Price = 9999
		p.Quantity = 9999
		result := db.Save(p)

		assert.NoError(t, result.Error)
		assert.Equal(t, int64(1), result.RowsAffected)

		assert.Equal(t, "Updated Test Product", p.Name)
		assert.Equal(t, uint(9999), p.Price)
		assert.Equal(t, uint(9999), p.Quantity)
	})

	t.Run("It shows Product", func(t *testing.T) {
		newProduct := &Product{
			Name: "Test Product",
			Price: 9999,
			Quantity: 9999,
		}
		db.Create(&newProduct)

		var p Product
		result := db.Last(&p)

		assert.NoError(t, result.Error)
		assert.Equal(t, int64(1), result.RowsAffected)

		assert.Equal(t, "Test Product", p.Name)
		assert.Equal(t, uint(9999), p.Price)
		assert.Equal(t, uint(9999), p.Quantity)
	})

	t.Run("It deletes Product", func(t *testing.T) {
		p := &Product{
			Name: "Test Product",
			Price: 9999,
			Quantity: 9999,
		}
		db.Create(&p)

		result := db.Delete(&Product{}, p.ID)

		assert.NoError(t, result.Error)
		assert.Equal(t, int64(1), result.RowsAffected)
	})
}

