package http

import (
	"log"
	"notik/internal/parts"
	"notik/internal/users/users_repo"
	"notik/pkg/httpErrors"
	"notik/pkg/logger"
	"notik/pkg/utils"

	"github.com/labstack/echo/v4"
)

type Handlers struct {
	log logger.Logger
	uc  parts.Usecase
}

func New(uc parts.Usecase, log logger.Logger) *Handlers {
	return &Handlers{uc: uc, log: log}
}

func (s *Handlers) Upsert() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := new(parts.UpsertInput)
		if err := utils.SanitizeRequest(c, input); err != nil {
			return utils.ErrResponseWithLog(c, s.log, err)
		}

		user, ok := c.Get("user").(users_repo.User)
		if !ok {
			return utils.ErrResponseWithLog(c, s.log, httpErrors.ErrUnauthorized)
		}

		part, err := s.uc.Upsert(c.Request().Context(), user.ID, *input)
		if err != nil {
			return utils.ErrResponseWithLog(c, s.log, err)
		}

		log.Print("after upsert")

		return c.JSON(200, part)
	}
}
