package controllers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vkazakevich/ebiznes/Go/dto"
	"github.com/vkazakevich/ebiznes/Go/models"
)

func (с *Controller) GetAllCartItems(ctx echo.Context) error {
	var cartItems []models.Cart
	с.DB.Preload("Product").Find(&cartItems)

	return ctx.JSON(http.StatusOK, cartItems)
}

func (с *Controller) CreateCartItem(ctx echo.Context) error {
	dto, err := contextToCartDto(ctx)
	if err != nil {
		return nil
	}

	p := models.Cart{ProductID: dto.ProductID, Quantity: dto.Quantity}
	с.DB.Create(&p)

	return ctx.JSON(http.StatusAccepted, p)
}

func (c *Controller) FindCartItem(ctx echo.Context) error {
	p, err := findCartItem(ctx, c)
	if err != nil {
		return nil
	}

	return ctx.JSON(http.StatusOK, p)
}

func (c *Controller) UpdateCartItem(ctx echo.Context) error {
	p, err := findCartItem(ctx, c)
	if err != nil {
		return nil
	}

	dto, err := contextToCartDto(ctx)
	if err != nil {
		return nil
	}

	p.ProductID = dto.ProductID
	p.Quantity = dto.Quantity

	c.DB.Save(&p)

	return ctx.JSON(http.StatusOK, p)
}

func (c *Controller) DeleteCartItem(ctx echo.Context) error {
	p, err := findCartItem(ctx, c)

	if err != nil {
		return nil
	}

	c.DB.Delete(p)
	return ctx.String(http.StatusNoContent, "")
}

func findCartItem(ctx echo.Context, c *Controller) (*models.Cart, error) {
	var cart models.Cart
	results := c.DB.Preload("Product").First(&cart, ctx.Param("id"))

	if results.Error != nil {
		_ = ctx.String(http.StatusNotFound, "not found")
		return nil, errors.New("cart item not found")
	}

	return &cart, nil
}

func contextToCartDto(ctx echo.Context) (*dto.CartDTO, error) {
	dto := new(dto.CartDTO)

	if err := ctx.Bind(dto); err != nil {
		_ = ctx.String(http.StatusBadRequest, "bad request")
		return nil, err
	}

	return dto, nil
}
