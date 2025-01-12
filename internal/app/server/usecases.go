package server

import (
	"github.com/arifinhermawan/amartha-billing-engine/internal/usecase/loan"
	"github.com/arifinhermawan/amartha-billing-engine/internal/usecase/user"
)

type UseCases struct {
	User *user.UseCase
	Loan *loan.UseCase
}

func NewUseCases(svc *Services) *UseCases {
	return &UseCases{
		User: user.NewUseCase(svc.Auth, svc.User),
		Loan: loan.NewUseCase(svc.Loan, svc.User),
	}
}
