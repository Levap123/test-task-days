package rest

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type DaysLeftBody struct {
	Days int `json:"days,omitempty"`
}

func (h *Handler) daysLeft(c echo.Context) error {
	daysLeft, err := h.service.Days.Left(time.Now())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, DaysLeftBody{Days: daysLeft})
	return nil
}
