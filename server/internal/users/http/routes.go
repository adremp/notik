package http

import (
	"notik/internal/users"
	"github.com/labstack/echo/v4"
)

func NewRoutes(g *echo.Group, h users.Handler) {
	g.POST("", h.Create())
}
