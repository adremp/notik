package http

import (
	"notik/internal/middleware"
	"notik/internal/users"

	"github.com/labstack/echo/v4"
)

func NewRoutes(g *echo.Group, h users.Handler, mv *middleware.MiddlewareManager) {

	g.POST("", h.Create())
}
