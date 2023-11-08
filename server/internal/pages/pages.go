package pages

import (
	"context"
	"notik/internal/pages/pages_repo"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	// GetByUserId() echo.HandlerFunc
	Create() echo.HandlerFunc
	// Delete() echo.HandlerFunc
	// Update() echo.HandlerFunc
}

type Usecase interface {
	// GetByUserId(ctx context.Context, userId string) ([]pages_repo.Page, error)
	Create(context.Context, pages_repo.CreateParams) (pages_repo.Page, error)
	// Delete(ctx context.Context, id string) error
	// Update(ctx context.Context, id string, title string) (pages_repo.Page, error)
}

type Repo interface {
	Create(context.Context, pages_repo.CreateParams) (pages_repo.Page, error)
}
