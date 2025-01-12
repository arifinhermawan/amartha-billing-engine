package server

import "github.com/arifinhermawan/amartha-billing-engine/internal/usecase/user"

type UseCases struct {
	User *user.UseCase
}

func NewUseCases(svc *Services) *UseCases {
	return &UseCases{
		User: user.NewUseCase(svc.Auth, svc.User),
	}
}
