package dto

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vkazakevich/ebiznes/Go/models"
)

func fakePaymentDto() (*PaymentDTO) {
	return &PaymentDTO{Amount: 10000000}
}

func TestPaymentDto (t *testing.T) {
	dto := fakePaymentDto()

    assert.NotNil(t, dto)

	assert.Equal(t, uint(10000000), dto.Amount)
}

func TestPaymentDto_Model(t *testing.T) {
	db := setupTestDB()
	dto := fakePaymentDto()

	p := models.Payment{Amount: dto.Amount}
	db.Create(&p)

	assert.Equal(t, p.Amount, dto.Amount)
	
	assert.NotNil(t, p.ID)
	assert.NotNil(t, p.CreatedAt)
	assert.NotNil(t, p.UpdatedAt)
}