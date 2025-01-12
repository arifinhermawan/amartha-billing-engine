package payment

import "context"

type paymentUseCaseProvider interface {
	PayWeeklyInstallment(ctx context.Context, loanID int64) error
}

type Handler struct {
	payment paymentUseCaseProvider
}

func NewHandler(payment paymentUseCaseProvider) *Handler {
	return &Handler{
		payment: payment,
	}
}
