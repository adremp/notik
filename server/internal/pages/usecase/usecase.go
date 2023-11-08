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

//	func (c *pagesUc) GetByUserId(ctx context.Context, userId string) ([]pages_repo.Page, error) {
//		_, err := strconv.Atoi(userId)
//		if err != nil {
//			return []pages_repo.Page{}, err
//		}
//		return nil, nil
//	}
func (c *pagesUc) Create(ctx context.Context, page pages_repo.CreateParams) (pages_repo.Page, error) {
	return c.db.Create(ctx, page)
}

// func (c *pagesUc) Update(ctx context.Context, id, title string) (pages_repo.Page, error) {
// 	return pages_repo.Page{}, nil
// }
// func (c *pagesUc) Delete(ctx context.Context, id string) error {
// 	return nil
// }
