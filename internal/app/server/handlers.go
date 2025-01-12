package server

import (
	"github.com/arifinhermawan/amartha-billing-engine/internal/handler/loan"
	"github.com/arifinhermawan/amartha-billing-engine/internal/handler/payment"
	"github.com/arifinhermawan/amartha-billing-engine/internal/handler/user"
)

type Handlers struct {
	Loan    *loan.Handler
	Payment *payment.Handler
	User    *user.Handler
}

func NewHandler(uc *UseCases) *Handlers {
	return &Handlers{
		Loan:    loan.NewHandler(uc.Loan),
		Payment: payment.NewHandler(uc.Payment),
		User:    user.NewHandler(uc.User),
	}
}
