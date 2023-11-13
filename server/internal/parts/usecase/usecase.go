package usecase

import (
	"context"
	"fmt"
	"notik/internal/pages"
	"notik/internal/pages/pages_repo"
	"notik/internal/parts"
	"notik/internal/parts/parts_repo"
	"notik/internal/users"
	"notik/pkg/httpErrors"
	"notik/pkg/logger"
)

type partsUc struct {
	db      parts.Repo
	usersUc users.Usecase
	pagesUc pages.Usecase
	log     logger.Logger
}

func New(db parts.Repo, usersUc users.Usecase, pagesUc pages.Usecase, log logger.Logger) *partsUc {
	return &partsUc{db: db, usersUc: usersUc, pagesUc: pagesUc, log: log}
}

func (s *partsUc) Upsert(ctx context.Context, userId int32, input parts.UpsertInput) (parts_repo.UpsertRow, error) {

	s.log.Infof("parts.uc.upsert: user_id: %d, input: %+v", userId, input)
	pages, err := s.pagesUc.GetByFields(ctx, pages_repo.GetByFieldsParams{ID: int32(input.PageID)})
	if len(pages) == 0 {
		return parts_repo.UpsertRow{}, httpErrors.NewNotFound("parts.uc.upsert: page not found")
	}

	if int32(pages[0].UserID) != userId {
		return parts_repo.UpsertRow{}, httpErrors.NewNotFound("parts.uc.upsert: page user_id != user_id")
	}

	part, err := s.db.Upsert(ctx, parts_repo.UpsertParams{
		Variant: parts_repo.PartType(input.Variant),
		Body:    input.Body,
		PageID:  input.PageID,
	})
	if err != nil {
		return parts_repo.UpsertRow{}, fmt.Errorf("parts.uc.upsert: %w", err)
	}

	return part, nil
}
