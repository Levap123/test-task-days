package rest

import (
	"log"
	"strings"

	"github.com/labstack/echo/v4"
)

const adminRole = "admin"

func (h *Handler) userRoleCheckMiddlware(next func(echo.Context) error) func(echo.Context) error {
	return func(c echo.Context) error {
		role :=strings.ToLower(c.Request().Header.Get("User-Role"))
		if strings.Contains(role, adminRole) {
			log.Println("red button user detected")
		}
		next(c)
		return nil
	}
}
