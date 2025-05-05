package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	// "github.com/vkazakevich/ebiznes/Go/models"
)

var (
	productJSON = `{"name": "Test Product", "price": 9999, "quantity": 1000}`
	updatedProductJSON = `{"name": "Updated Test Product", "price": 100, "quantity": 5}`
)

func TestCreateProduct(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(productJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	controller := &Controller{DB: mockDB}

	if assert.NoError(t, controller.CreateProduct(c)) {
		var resp map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusAccepted, rec.Code)
		assert.Equal(t, "Test Product", resp["name"])
	}
}

func TestGetAllProduct(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	controller := &Controller{DB: mockDB}

	if assert.NoError(t, controller.GetAllProduct(c)) {
		var resp []map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Greater(t, len(resp), 0)
	}
}

func TestFindProduct(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	controller := &Controller{DB: mockDB}

	if assert.NoError(t, controller.FindProduct(c)) {
		var resp map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Greater(t, resp["ID"], float64(0))
	}
}

func TestFindProduct_NotFoundError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("99")
	controller := &Controller{DB: mockDB}

	if assert.NoError(t, controller.FindProduct(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
}

func TestUpdateProduct(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/products", strings.NewReader(updatedProductJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	controller := &Controller{DB: mockDB}

	if assert.NoError(t, controller.UpdateProduct(c)) {
		var resp map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Updated Test Product", resp["name"])
	}
}

func TestDeleteProduct(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/products", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	controller := &Controller{DB: mockDB}

	if assert.NoError(t, controller.DeleteProduct(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}