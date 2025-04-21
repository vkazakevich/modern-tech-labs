package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vkazakevich/ebiznes/Go/dto"
	"github.com/vkazakevich/ebiznes/Go/models"
)

func (с *Controller) CreatePayment(ctx echo.Context) error {
	dto := new(dto.PaymentDTO)

	if err := ctx.Bind(dto); err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	p := models.Payment{Amount: dto.Amount}
	с.DB.Create(&p)

	return ctx.JSON(http.StatusAccepted, p)
}

