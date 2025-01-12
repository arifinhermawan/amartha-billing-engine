package loan

import (
	"context"

	"github.com/arifinhermawan/amartha-billing-engine/internal/service/loan"
	"github.com/arifinhermawan/amartha-billing-engine/internal/service/user"
)

type loanServiceProvider interface {
	CreateLoan(ctx context.Context, req loan.CreateLoanReq) error
}

type userServiceProvider interface {
	GetUserByID(ctx context.Context, userID int64) (user.User, error)
}

type UseCase struct {
	loan loanServiceProvider
	user userServiceProvider
}

func NewUseCase(loan loanServiceProvider, user userServiceProvider) *UseCase {
	return &UseCase{
		loan: loan,
		user: user,
	}
}
