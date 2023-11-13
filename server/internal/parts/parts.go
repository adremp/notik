package parts

import (
	"context"
	"notik/internal/parts/parts_repo"

	"github.com/labstack/echo/v4"
)

type Handlers interface {
	Upsert() echo.HandlerFunc
}

type Usecase interface {
	Upsert(ctx context.Context, userId int32, input UpsertInput) (parts_repo.UpsertRow, error)
}

type Repo interface {
	Upsert(ctx context.Context, input parts_repo.UpsertParams) (parts_repo.UpsertRow, error)
	// GetByFields(ctx context.Context, input parts_repo.UpsertParams) ([]parts_repo.Part, error)
}

type UpsertInput struct {
	Variant string `json:"variant" validate:"required,oneof=header1 header2 header3 text image"`
	Body    string `json:"body" validate:"required"`
	PageID  int64  `json:"page_id" validate:"required"`
}
