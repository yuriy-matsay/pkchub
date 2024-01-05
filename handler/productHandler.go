package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetCategories(c echo.Context) error {
	categories, err := h.services.Storage.GetCategories()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.Render(http.StatusOK, "startpage", categories)
}

func (h *Handler) GetGoodsByCategory(c echo.Context) error {
	id := c.Param("id")

	products, err := h.services.Storage.GetGoodsByCategory(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	brands, err := h.services.Storage.GetBrandsByCategory(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	var data = map[string]interface{}{}
	data["Products"] = products
	data["Brands"] = brands
	data["Category"] = id

	return c.Render(http.StatusOK, "categoryitems", data)
}

func (h *Handler) GetGoodsByBrand(c echo.Context) error {
	id := c.Param("id")
	category := c.QueryParam("category")

	products, err := h.services.Storage.GetGoodsByBrand(id, category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	var data = map[string]interface{}{}
	data["Products"] = products

	return c.Render(http.StatusOK, "branditems", data)
}
