package loan

import (
	"context"

	"github.com/arifinhermawan/amartha-billing-engine/internal/usecase/loan"
)

type loanUseCaseProvider interface {
	CreateLoan(ctx context.Context, req loan.CreateLoanReq) error
	GetOutstandingBalance(ctx context.Context, loanID int64) (float64, error)
}

type Handler struct {
	loan loanUseCaseProvider
}

func NewHandler(loan loanUseCaseProvider) *Handler {
	return &Handler{
		loan: loan,
	}
}
