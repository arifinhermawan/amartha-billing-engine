package user

import "context"

type authServiceProvider interface {
	EncryptPassword(password string) string
}

type userServiceProvider interface {
	CreateUser(ctx context.Context, name string, hashedPassword string) error
}

type UseCase struct {
	auth authServiceProvider
	user userServiceProvider
}

func NewUseCase(auth authServiceProvider, user userServiceProvider) *UseCase {
	return &UseCase{
		auth: auth,
		user: user,
	}
}
