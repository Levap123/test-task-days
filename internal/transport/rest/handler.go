package rest

import (
	"github.com/Levap123/test-task-days/internal/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *echo.Echo {
	r := echo.New()
	r.GET("/daysleft", h.userRoleCheckMiddlware(h.daysLeft))
	return r
}
