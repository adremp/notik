package pages

import (
	"context"
	"notik/internal/pages/pages_repo"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	Create() echo.HandlerFunc
}

type Usecase interface {
	Create(context.Context, pages_repo.CreateParams) (pages_repo.CreateRow, error)
	GetByFields(context.Context, pages_repo.GetByFieldsParams) ([]pages_repo.GetByFieldsRow, error)
}

type Repo interface {
	Create(context.Context, pages_repo.CreateParams) (pages_repo.CreateRow, error)
	GetByFields(context.Context, pages_repo.GetByFieldsParams) ([]pages_repo.GetByFieldsRow, error)
}

type CreateInput struct {
	Title string `validate:"required,min=3,max=80"`
}
