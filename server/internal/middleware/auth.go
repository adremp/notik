package middleware

import (
	"notik/internal/users/users_repo"
	"notik/pkg/httpErrors"
	"notik/pkg/utils"
	"strings"

	"github.com/labstack/echo/v4"
)

func (s *MiddlewareManager) AuthJWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		tokenCookie, err := c.Cookie("token")
		if err != nil || strings.TrimSpace(tokenCookie.Value) == "" {
			return c.JSON(httpErrors.RequestError(httpErrors.ErrUnauthorized))
		}

		userClaims, err := utils.ParseToken(tokenCookie.Value)
		if err != nil {
			return c.JSON(httpErrors.RequestError(httpErrors.ErrUnauthorized))
		}

		users, err := s.usersUc.GetByFields(c.Request().Context(), users_repo.GetByFieldsParams{ID: userClaims.Id})

		if len(users) == 0 || err != nil {
			return c.JSON(httpErrors.RequestError(httpErrors.ErrUnauthorized))
		}

		c.Set("user", users[0])

		return next(c)
	}
}
