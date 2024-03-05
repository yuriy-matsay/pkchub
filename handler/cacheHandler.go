package handler

import (
	"encoding/json"
	"log"
	"pkhub/models"
)

const (
	curr string = "currencies"
	cat  string = "categories"
)

func (h *Handler) getCurrencies() (currencies string) {
	currencies, err := h.services.Cache.Get(curr)
	if err != nil {
		log.Print("key ", curr, " ", err)

		currencies = h.services.Storage.GetCurrencies()
		go h.services.Cache.Set(curr, currencies)

		return
	}
	return
}

func (h *Handler) getCategories() (categories []models.Category, err error) {
	val, err := h.services.Cache.Get(cat)
	if err != nil {
		log.Print("key ", cat, " ", err)

		categories, err = h.services.Storage.GetCategories()
		if err != nil {
			log.Print(err)
			return
		}
		go h.cacheJson(cat, categories)

		return categories, nil
	}
	err = json.Unmarshal([]byte(val), &categories)
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (h *Handler) cacheJson(key string, value interface{}) {
	jsonData, err := json.Marshal(value)
	if err != nil {
		log.Print(err)
	}
	h.services.Cache.Set(cat, jsonData)
}
