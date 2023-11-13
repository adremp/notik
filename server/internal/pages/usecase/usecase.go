package usecase

import (
	"context"
	"notik/internal/pages"
	"notik/internal/pages/pages_repo"
)

type pagesUc struct {
	db pages.Repo
}

func New(repo pages.Repo) pages.Usecase {
	return &pagesUc{repo}
}

func (s *pagesUc) Create(ctx context.Context, page pages_repo.CreateParams) (pages_repo.CreateRow, error) {
	return s.db.Create(ctx, page)
}

func (s *pagesUc) GetByFields(ctx context.Context, fields pages_repo.GetByFieldsParams) ([]pages_repo.GetByFieldsRow, error) {
	return s.db.GetByFields(ctx, fields)
}
