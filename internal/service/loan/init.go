package loan

import (
	"context"
	"database/sql"
	"time"

	"github.com/arifinhermawan/amartha-billing-engine/internal/lib/configuration"
	"github.com/arifinhermawan/amartha-billing-engine/internal/repository/pgsql"
)

type libProvider interface {
	GetConfig() *configuration.AppConfig
	GetTimeGMT7() time.Time
}

type dbProvider interface {
	BeginTX(ctx context.Context, options *sql.TxOptions) (*sql.Tx, error)
	CreateLoanInDB(ctx context.Context, tx *sql.Tx, req pgsql.CreateLoanReq) (int64, error)
	CreatePaymentPlanInDB(ctx context.Context, tx *sql.Tx, plans []pgsql.PaymentPlan) error
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
