package user

import (
	"context"
	"time"

	"github.com/arifinhermawan/amartha-billing-engine/internal/service/user"
)

type libProvider interface {
	// GetTimeGMT7 retrieves the current time in the GMT+7.
	GetTimeGMT7() time.Time
}

type authServiceProvider interface {
	EncryptPassword(password string) string
}

type userServiceProvider interface {
	CreateUser(ctx context.Context, name string, hashedPassword string) error
	GetDelinquentUsers(ctx context.Context, date time.Time) ([]user.User, error)
}

type UseCase struct {
	lib  libProvider
	auth authServiceProvider
	user userServiceProvider
}

func NewUseCase(lib libProvider, auth authServiceProvider, user userServiceProvider) *UseCase {
	return &UseCase{
		lib:  lib,
		auth: auth,
		user: user,
	}
}
