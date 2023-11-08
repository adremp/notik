package http

import (
	"notik/internal/pages"

	"github.com/labstack/echo/v4"
)


func NewRoutes(c *echo.Group, h pages.Handler) {
	c.POST("", h.Create())
}