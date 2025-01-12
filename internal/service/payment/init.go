package payment

import (
	"context"
	"time"

	"github.com/arifinhermawan/amartha-billing-engine/internal/repository/pgsql"
)

type libProvider interface {
	GetTimeGMT7() time.Time
}

type dbProvider interface {
	GetUnpaidPaymentPlanByLoanID(ctx context.Context, loanID int64) ([]pgsql.PaymentPlan, error)
	UpdatePaymentPlan(ctx context.Context, req pgsql.UpdatePaymentPlanReq) error
}

type Service struct {
	lib libProvider
	db  dbProvider
}

func NewService(lib libProvider, db dbProvider) *Service {
	return &Service{
		lib: lib,
		db:  db,
	}
}
