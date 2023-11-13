package users

import (
	"context"
	"notik/internal/users/users_repo"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Handler interface {
	Create() echo.HandlerFunc
}

type Repo interface {
	Create(context.Context, users_repo.CreateParams) (users_repo.CreateRow, error)
	GetByFields(context.Context, users_repo.GetByFieldsParams) ([]users_repo.User, error)
}

type Usecase interface {
	Create(context.Context, CreateInput) (*UserWithToken, error)
	GetByFields(context.Context, users_repo.GetByFieldsParams) ([]users_repo.User, error)
}

type CreateInput struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" validate:"required,min=8,max=20"`
}

func (s *CreateInput) HashPassword() error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(s.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	s.Password = string(hashed)
	return nil
}

func (s *CreateInput) ComparePassword(pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(s.Password), []byte(pass))
}

type UserWithToken struct {
	User  users_repo.CreateRow `json:"user"`
	Token string               `json:"token"`
}
