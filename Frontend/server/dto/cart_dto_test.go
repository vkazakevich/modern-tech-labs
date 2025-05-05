package dto

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vkazakevich/ebiznes/Go/models"
)

func fakeCartDto() (*CartDTO) {
	return &CartDTO{
		ProductID: 1000,
		Quantity: 9999,
	}
}

func TestCartDto (t *testing.T) {
	dto := fakeCartDto()

    assert.NotNil(t, dto)

	assert.Equal(t, uint(1000), dto.ProductID)
	assert.Equal(t, uint(9999), dto.Quantity)
}

func TestCartDto_Model(t *testing.T) {
	db := setupTestDB()
	dto := fakeCartDto()

	c := models.Cart{ProductID: dto.ProductID, Quantity: dto.Quantity}
	db.Create(&c)

	assert.Equal(t, c.ProductID, dto.ProductID)
	assert.Equal(t, c.Quantity, dto.Quantity)

	assert.NotNil(t, c.ID)
	assert.NotNil(t, c.CreatedAt)
	assert.NotNil(t, c.UpdatedAt)
}