package http

import (
	"net/http"
	"notik/internal/pages"
	"notik/internal/pages/pages_repo"
	"notik/pkg/httpErrors"
	"notik/pkg/utils"

	"github.com/labstack/echo/v4"
)

type pagesHandler struct {
	uc pages.Usecase
}

func New(uc pages.Usecase) pages.Handler {
	return &pagesHandler{uc: uc}
}

func (s *pagesHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input pages.CreateInput
		if err := utils.SanitizeRequest(c, &input); err != nil {
			return c.JSON(httpErrors.RequestError(err))
		}

		page, err := s.uc.Create(c.Request().Context(), pages_repo.CreateParams{Title: input.Title, UserID: 1})
		if err != nil {
			return c.JSON(httpErrors.RequestError(err))
		}

		return c.JSON(http.StatusOK, page)
	}
}
