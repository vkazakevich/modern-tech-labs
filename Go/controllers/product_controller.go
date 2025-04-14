package controllers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vkazakevich/ebiznes/Go/dto"
	"github.com/vkazakevich/ebiznes/Go/models"
)

func (с *Controller) GetAllProduct(ctx echo.Context) error {
	var products []models.Product
	с.DB.Find(&products)

	return ctx.JSON(http.StatusOK, products)
}

func (с *Controller) CreateProduct(ctx echo.Context) error {
	dto, err := contextToProductDto(ctx)
	if err != nil {
		return nil
	}

	p := models.Product{Name: dto.Name, Quantity: dto.Quantity, Price: dto.Price}
	с.DB.Create(&p)

	return ctx.JSON(http.StatusAccepted, p)
}

func (c *Controller) FindProduct(ctx echo.Context) error {
	p, err := findProduct(ctx, c)
	if err != nil {
		return nil
	}

	return ctx.JSON(http.StatusOK, p)
}

func (c *Controller) UpdateProduct(ctx echo.Context) error {
	p, err := findProduct(ctx, c)
	if err != nil {
		return nil
	}

	dto, err := contextToProductDto(ctx)
	if err != nil {
		return nil
	}

	p.Name = dto.Name
	p.Quantity = dto.Quantity
	p.Price = dto.Price

	c.DB.Save(&p)

	return ctx.JSON(http.StatusOK, p)
}

func (c *Controller) DeleteProduct(ctx echo.Context) error {
	p, err := findProduct(ctx, c)

	if err != nil {
		return nil
	}

	c.DB.Delete(p)
	return ctx.String(http.StatusNoContent, "")
}

func findProduct(ctx echo.Context, c *Controller) (*models.Product, error) {
	var product models.Product
	results := c.DB.First(&product, ctx.Param("id"))

	if results.Error != nil {
		_ = ctx.String(http.StatusNotFound, "not found")
		return nil, errors.New("product not found")
	}

	return &product, nil
}

func contextToProductDto(ctx echo.Context) (*dto.ProductDTO, error) {
	dto := new(dto.ProductDTO)

	if err := ctx.Bind(dto); err != nil {
		_ = ctx.String(http.StatusBadRequest, "bad request")
		return nil, err
	}

	return dto, nil
}
