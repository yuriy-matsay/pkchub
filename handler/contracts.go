package handler

import (
	"encoding/json"
	"pkhub/service"
)

const (
	curr string = "currencies"
	cat  string = "categories"
)

type Handler struct {
	services   *service.Service
	currencies string
}

func NewHandler(s *service.Service) *Handler {
	value := s.Storage.GetCurrencies()
	s.Cache.Set(curr, value)

	categories, err := s.Storage.GetCategories()
	if err != nil {
		panic(err)
	}
	jsonData, err := json.Marshal(categories)
	if err != nil {
		panic(err)
	}
	s.Cache.Set(cat, jsonData)

	return &Handler{
		services:   s,
		currencies: s.Cache.Get(curr),
	}
}
