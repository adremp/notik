package http

import (
	"notik/internal/middleware"
	"notik/internal/parts"

	"github.com/labstack/echo/v4"
)

func NewRoutes(g *echo.Group, h parts.Handlers, mv *middleware.MiddlewareManager) {
	g.Use(mv.AuthJWTMiddleware)
	g.POST("", h.Upsert())
}
