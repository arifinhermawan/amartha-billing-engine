package user

import (
	"context"
	"time"

	"github.com/arifinhermawan/amartha-billing-engine/internal/usecase/user"
)

type userUseCaseProvider interface {
	CreateUser(ctx context.Context, name string, password string) error
	GetDelinquentsUsers(ctx context.Context, date time.Time) ([]user.User, error)
}

type Handler struct {
	user userUseCaseProvider
}

func NewHandler(user userUseCaseProvider) *Handler {
	return &Handler{
		user: user,
	}
}
