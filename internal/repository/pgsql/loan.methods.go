package pgsql

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

func (r *Repository) CreateLoanInDB(ctx context.Context, tx *sql.Tx, req CreateLoanReq) (int64, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.lib.GetConfig().Database.DefaultTimeout)*time.Second)
	defer cancel()

	metadata := map[string]interface{}{
		"user_id":             req.UserID,
		"amount":              req.Amount,
		"interest_rate":       req.InterestRate,
		"outstanding_balance": req.OutstandingBalance,
	}

	namedQuery, args, err := sqlx.Named(queryCreateLoan, req)
	if err != nil {
		log.Printf("[CreateLoanInDB] sqlxNamed() got error: %v\nMetadata: %v\n", err, metadata)
		return 0, err
	}

	var id int64
	err = tx.QueryRowContext(ctxTimeout, r.db.Rebind(namedQuery), args...).Scan(&id)
	if err != nil {
		log.Printf("[CreateLoanInDB] tx.QueryRowContext() got error: %v\nMetadata: %v\n", err, metadata)
		return 0, err
	}

	return id, nil
}
