package loan

import (
	"context"
	"database/sql"
	"time"

	"github.com/arifinhermawan/amartha-billing-engine/internal/repository/pgsql"
)

type libProvider interface {
	GetTimeGMT7() time.Time
}

type dbProvider interface {
	BeginTX(ctx context.Context, options *sql.TxOptions) (*sql.Tx, error)
	CreateLoanInDB(ctx context.Context, tx *sql.Tx, req pgsql.CreateLoanReq) (int64, error)
	CreatePaymentPlanInDB(ctx context.Context, tx *sql.Tx, plans []pgsql.PaymentPlan) error
	GetLoanByIDFromDB(ctx context.Context, loanID int64) (pgsql.Loan, error)
	GetOutstandingBalanceFromDB(ctx context.Context, loanID int64) (float64, error)
	UpdateLoanInDB(ctx context.Context, req pgsql.UpdateLoanReq) error
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
