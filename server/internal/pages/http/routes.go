package http

import (
	"notik/internal/middleware"
	"notik/internal/pages"

	"github.com/labstack/echo/v4"
)

func NewRoutes(g *echo.Group, h pages.Handler, mv *middleware.MiddlewareManager) {
	g.Use(mv.AuthJWTMiddleware)
	g.POST("", h.Create())
}
