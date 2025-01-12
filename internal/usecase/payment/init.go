package payment

import (
	"context"

	"github.com/arifinhermawan/amartha-billing-engine/internal/service/loan"
)

type loanServiceProvider interface {
	GetLoanByID(ctx context.Context, loanID int64) (loan.Loan, error)
	UpdateLoan(ctx context.Context, loanID int64, outstandingBalance float64) error
}

type paymentServiceProvider interface {
	PayWeeklyInstallment(ctx context.Context, loanID int64) (float64, error)
}

type UseCase struct {
	loan    loanServiceProvider
	payment paymentServiceProvider
}

func NewUseCase(loan loanServiceProvider, payment paymentServiceProvider) *UseCase {
	return &UseCase{
		loan:    loan,
		payment: payment,
	}
}
