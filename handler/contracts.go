package handler

import (
	"pkhub/service"
)

type Handler struct {
	services   *service.Service
	currencies string
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		services:   s,
		currencies: s.Storage.GetCurrencies(),
	}
}
