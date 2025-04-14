package controllers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vkazakevich/ebiznes/Go/dto"
	"github.com/vkazakevich/ebiznes/Go/models"
)

func (с *Controller) GetAllCategories(ctx echo.Context) error {
	var categories []models.Category

	if ctx.QueryParam("hasProducts") == "1" {
		с.DB.Scopes(models.WithProducts, models.HasProducts).Find(&categories)
	} else {
		с.DB.Scopes(models.WithProducts).Find(&categories)
	}

	return ctx.JSON(http.StatusOK, categories)
}

func (с *Controller) CreateCategory(ctx echo.Context) error {
	dto, err := contextToCategoryDto(ctx)
	if err != nil {
		return nil
	}

	cat := models.Category{Name: dto.Name}
	с.DB.Create(&cat)

	return ctx.JSON(http.StatusAccepted, cat)
}

func (c *Controller) FindCategory(ctx echo.Context) error {
	cat, err := findCategory(ctx, c)
	if err != nil {
		return nil
	}

	return ctx.JSON(http.StatusOK, cat)
}

func (c *Controller) UpdateCategory(ctx echo.Context) error {
	cat, err := findCategory(ctx, c)
	if err != nil {
		return nil
	}

	dto, err := contextToCategoryDto(ctx)
	if err != nil {
		return nil
	}

	cat.Name = dto.Name
	c.DB.Save(&cat)

	return ctx.JSON(http.StatusOK, cat)
}

func (c *Controller) DeleteCategory(ctx echo.Context) error {
	cat, err := findCategory(ctx, c)

	if err != nil {
		return nil
	}

	c.DB.Delete(cat)
	return ctx.String(http.StatusNoContent, "")
}

func findCategory(ctx echo.Context, c *Controller) (*models.Category, error) {
	var cat models.Category
	results := c.DB.Scopes(models.WithProducts).First(&cat, ctx.Param("id"))

	if results.Error != nil {
		_ = ctx.String(http.StatusNotFound, "not found")
		return nil, errors.New("category not found")
	}

	return &cat, nil
}

func contextToCategoryDto(ctx echo.Context) (*dto.CategoryDTO, error) {
	dto := new(dto.CategoryDTO)

	if err := ctx.Bind(dto); err != nil {
		_ = ctx.String(http.StatusBadRequest, "bad request")
		return nil, err
	}

	return dto, nil
}
