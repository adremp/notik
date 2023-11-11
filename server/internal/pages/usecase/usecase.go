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

func (s *pagesUc) Create(ctx context.Context, page pages_repo.CreateParams) (pages_repo.Page, error) {
	return s.db.Create(ctx, page)
}
