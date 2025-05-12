package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var paymentJSON = `{"amount": 1000}`

func TestCreatePayment(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/payments", strings.NewReader(paymentJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	controller := &Controller{DB: mockDB}

	if assert.NoError(t, controller.CreatePayment(c)) {
		var resp map[string]interface{}
		err := json.Unmarshal(rec.Body.Bytes(), &resp)

		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusAccepted, rec.Code)
			assert.Equal(t, float64(1000), resp["Amount"])
		}
	}
}

func TestCreatePaymentWithError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/payments", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	controller := &Controller{DB: mockDB}

	if assert.NoError(t, controller.CreatePayment(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}