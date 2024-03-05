package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pkhub/models"
)

func (h *Handler) GetCategories(c echo.Context) error {
	var categories []models.Category
	categories, err := h.getCategories()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	var data = map[string]interface{}{}
	data["Categories"] = categories
	data["Currencies"] = h.getCurrencies()

	return c.Render(http.StatusOK, "startpage", data)
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
	data["Currencies"] = h.getCurrencies()

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
	data["Currencies"] = h.getCurrencies()

	return c.Render(http.StatusOK, "branditems", data)
}

func (h *Handler) GetItem(c echo.Context) error {
	id := c.Param("id")

	params := h.services.Storage.GetItemParams(id)

	var data = map[string]interface{}{}
	data["Currencies"] = h.getCurrencies()
	data["Params"] = params
	data["Info"] = h.services.Storage.GetItemInfo(id)

	return c.Render(http.StatusOK, "item", data)
}

func (h *Handler) Update(c echo.Context) error {
	currensies := h.services.Storage.GetCurrencies()
	h.services.Cache.Set(curr, currensies)

	return c.HTML(http.StatusOK, "updates")
}
