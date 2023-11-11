package usecase

import (
	"context"
	"fmt"
	"notik/internal/users"
	"notik/internal/users/users_repo"
	"notik/pkg/httpErrors"
	"notik/pkg/utils"
	"time"
)

type usersUc struct {
	repo users.Repo
}

func New(repo users.Repo) users.Usecase {
	return &usersUc{repo}
}

func (s *usersUc) Create(ctx context.Context, input users.CreateInput) (*users.UserWithToken, error) {
	userList, err := s.repo.GetByFields(ctx, users_repo.GetByFieldsParams{Email: input.Email})
	if len(userList) != 0 || err == nil {
		return nil, httpErrors.ErrEmailExist
	}

	if err := input.HashPassword(); err != nil {
		return nil, fmt.Errorf("users.uc.create.hashPassword: %w", err)
	}

	newUser, err := s.repo.Create(ctx, users_repo.CreateParams{Username: input.Username, Email: input.Email, Password: input.Password})
	if err != nil {
		return nil, fmt.Errorf("users.uc.create: %w", err)
	}

	token, err := utils.GenerateToken(newUser.ID, newUser.Email, time.Minute*10)
	if err != nil {
		return nil, fmt.Errorf("users.uc.create.generateToken: %w", err)
	}

	return &users.UserWithToken{User: newUser, Token: token}, nil
}

func (s *usersUc) GetByFields(ctx context.Context, fields users_repo.GetByFieldsParams) ([]users_repo.User, error) {
	users, err := s.repo.GetByFields(ctx, fields)
	if err != nil {
		return nil, fmt.Errorf("users.uc.GetByFields: %w", err)
	}
	return users, nil
}
