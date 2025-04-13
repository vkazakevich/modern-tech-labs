package controllers

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ProductDTO struct {
	Name     string `json:"name" form:"name" query:"name"`
	Quantity int    `json:"quantity" form:"quantity" query:"quantity"`
}

type Product struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Quantity int       `json:"quantity"`
}

var products []Product

func findProductByParam(c echo.Context) (*Product, error) {
	id := c.Param("id")
	uuid, err := uuid.Parse(id)

	if err != nil {
		_ = c.String(http.StatusBadRequest, "bad request")
		return nil, err
	}

	for i := range products {
		if products[i].ID == uuid {
			return &products[i], nil
		}
	}

	_ = c.String(http.StatusNotFound, "not found")
	return nil, errors.New("Product not found")
}

func contextToProductDto(c echo.Context) (*ProductDTO, error) {
	dto := new(ProductDTO)

	if err := c.Bind(dto); err != nil {
		_ = c.String(http.StatusBadRequest, "bad request")
		return nil, err
	}

	return dto, nil
}

func GetAllProduct(c echo.Context) error {
	if products == nil {
		return c.JSON(http.StatusOK, []Product{})
	}

	return c.JSON(http.StatusOK, products)
}

func CreateProduct(c echo.Context) error {
	dto, err := contextToProductDto(c)
	if err != nil {
		return nil
	}

	p := Product{ID: uuid.New(), Name: dto.Name, Quantity: dto.Quantity}
	products = append(products, p)
	return c.JSON(http.StatusAccepted, p)
}

func FindProduct(c echo.Context) error {
	p, err := findProductByParam(c)
	if err != nil {
		return nil
	}

	return c.JSON(http.StatusOK, p)
}

func UpdateProduct(c echo.Context) error {
	p, err := findProductByParam(c)
	if err != nil {
		return nil
	}

	dto, err := contextToProductDto(c)
	if err != nil {
		return nil
	}

	p.Name = dto.Name
	p.Quantity = dto.Quantity

	return c.JSON(http.StatusOK, p)
}

func DeleteProduct(c echo.Context) error {
	product, err := findProductByParam(c)
	if err != nil {
		return nil
	}

	for i, p := range products {
		if p.ID == product.ID {
			products = append(products[:i], products[i+1:]...)
		}
	}

	return c.String(http.StatusNoContent, "")
}
