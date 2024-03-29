package handler

import (
	"pkhub/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler {

	return &Handler{
		services: s,
	}
}
