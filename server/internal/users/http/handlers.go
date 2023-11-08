package http

import (
	"log"
	"net/http"
	"notik/internal/users"
	"notik/pkg/httpErrors"
	"notik/pkg/utils"

	"github.com/labstack/echo/v4"
)

type usersHandler struct {
	uc users.Usecase
}

func New(uc users.Usecase) users.Handler {
	return &usersHandler{uc}
}

func (s *usersHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input users.CreateInput

		if err := utils.SanitizeRequest(c, &input); err != nil {
			log.Print(err)
			return c.JSON(httpErrors.RequestError(err))
		}

		createdUser, err := s.uc.Create(c.Request().Context(), input)
		if err != nil {
			log.Print(err)
			return c.JSON(httpErrors.RequestError(err))
		}

		c.SetCookie(utils.CreateCookie(createdUser.Token, 600))

		return c.JSON(http.StatusCreated, createdUser.User)
	}
}
