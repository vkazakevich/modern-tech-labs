package dto

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vkazakevich/ebiznes/Go/models"
)

func fakeProductDto() (*ProductDTO) {
	return &ProductDTO{
		Name: "Test Product",
		Quantity: 9999,
		Price: 9999,
		CategoryID: 1,
	}
}

func TestProductDto (t *testing.T) {
	dto := fakeProductDto()

    assert.NotNil(t, dto)

	assert.Equal(t, "Test Product", dto.Name)
	assert.Equal(t, uint(9999), dto.Quantity)
	assert.Equal(t, uint(9999), dto.Price)
	assert.Equal(t, uint(1), dto.CategoryID)
}

func TestProductDto_Model(t *testing.T) {
	dto := fakeProductDto()
	db := setupTestDB()

	p := models.Product{
		Name: dto.Name, 
		Quantity: dto.Quantity, 
		Price: dto.Price,
		CategoryID: dto.CategoryID,
	}
	db.Create(&p)

	assert.Equal(t, p.Name, dto.Name)
	assert.Equal(t, p.Quantity, dto.Quantity)
	assert.Equal(t, p.Price, dto.Price)
	assert.Equal(t, p.CategoryID, dto.CategoryID)
	
	assert.NotNil(t, p.ID)
	assert.NotNil(t, p.CreatedAt)
	assert.NotNil(t, p.UpdatedAt)
}