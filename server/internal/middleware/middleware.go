package middleware

import "notik/internal/users"

type MiddlewareManager struct {
	usersUc users.Usecase
}

func New(usersUc users.Usecase) *MiddlewareManager {
	return &MiddlewareManager{usersUc}
}
