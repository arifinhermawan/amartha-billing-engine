package server

import (
	"github.com/arifinhermawan/amartha-billing-engine/internal/handler/loan"
	"github.com/arifinhermawan/amartha-billing-engine/internal/handler/user"
)

type Handlers struct {
	User *user.Handler
	Loan *loan.Handler
}

func NewHandler(uc *UseCases) *Handlers {
	return &Handlers{
		User: user.NewHandler(uc.User),
		Loan: loan.NewHandler(uc.Loan),
	}
}
