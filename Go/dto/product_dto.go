package dto

type ProductDTO struct {
	Name     string `json:"name" form:"name" query:"name"`
	Quantity uint   `json:"quantity" form:"quantity" query:"quantity"`
	Price    uint   `json:"price" form:"price" query:"price"`
}
