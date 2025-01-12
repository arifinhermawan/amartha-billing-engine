package server

import (
	"github.com/arifinhermawan/amartha-billing-engine/internal/usecase/loan"
	"github.com/arifinhermawan/amartha-billing-engine/internal/usecase/payment"
	"github.com/arifinhermawan/amartha-billing-engine/internal/usecase/user"
)

type UseCases struct {
	Loan    *loan.UseCase
	Payment *payment.UseCase
	User    *user.UseCase
}

func NewUseCases(svc *Services) *UseCases {
	return &UseCases{
		Loan:    loan.NewUseCase(svc.Loan, svc.User),
		Payment: payment.NewUseCase(svc.Loan, svc.Payment),
		User:    user.NewUseCase(svc.Auth, svc.User),
	}
}
