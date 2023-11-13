package http

import (
	"net/http"
	"notik/internal/users"
	"notik/pkg/logger"
	"notik/pkg/utils"

	"github.com/labstack/echo/v4"
)

type usersHandler struct {
	uc  users.Usecase
	log logger.Logger
}

func New(uc users.Usecase, log logger.Logger) users.Handler {
	return &usersHandler{uc, log}
}

func (s *usersHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input users.CreateInput

		if err := utils.SanitizeRequest(c, &input); err != nil {
			return utils.ErrResponseWithLog(c, s.log, err)
		}

		createdUser, err := s.uc.Create(c.Request().Context(), input)
		if err != nil {
			return utils.ErrResponseWithLog(c, s.log, err)
		}

		c.SetCookie(utils.CreateCookie(createdUser.Token, 600))

		return c.JSON(http.StatusCreated, createdUser.User)
	}
}
