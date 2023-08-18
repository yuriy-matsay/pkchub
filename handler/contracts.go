package handler

import (
	// "github.com/labstack/echo/v4"
	"pkhub/service"
)

// type HandlerInterface interface {
// 	GetCategories(c echo.Context) error
// 	GetProducts(c echo.Context) error
// }

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		services: s,
	}
}
