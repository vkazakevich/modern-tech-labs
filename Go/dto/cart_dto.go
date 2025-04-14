package dto

type CartDTO struct {
	ProductID uint `json:"product_id" form:"product_id" query:"product_id"`
	Quantity  uint `json:"quantity" form:"quantity" query:"quantity"`
}
