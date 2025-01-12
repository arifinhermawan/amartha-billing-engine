package loan

import (
	"context"

	"github.com/arifinhermawan/amartha-billing-engine/internal/usecase/loan"
)

type loanUseCaseProvider interface {
	CreateLoan(ctx context.Context, req loan.CreateLoanReq) error
}

type Handler struct {
	loan loanUseCaseProvider
}

func NewHandler(loan loanUseCaseProvider) *Handler {
	return &Handler{
		loan: loan,
	}
}
