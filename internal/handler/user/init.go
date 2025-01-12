package user

import "context"

type userUseCaseProvider interface {
	CreateUser(ctx context.Context, name string, password string) error
}

type Handler struct {
	user userUseCaseProvider
}

func NewHandler(user userUseCaseProvider) *Handler {
	return &Handler{
		user: user,
	}
}
