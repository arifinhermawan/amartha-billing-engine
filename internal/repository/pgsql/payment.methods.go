package pgsql

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

func (r *Repository) CreatePaymentPlanInDB(ctx context.Context, tx *sql.Tx, plans []PaymentPlan) error {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.lib.GetConfig().Database.DefaultTimeout)*time.Second)
	defer cancel()

	params := make([]map[string]interface{}, len(plans))
	loanID := int64(0)
	for idx, plan := range plans {
		loanID = plan.LoanID
		params[idx] = map[string]interface{}{
			"loan_id":     loanID,
			"week_number": plan.WeekNumber,
			"amount":      plan.Amount,
			"due_date":    plan.DueDate,
			"created_at":  plan.CreatedAt,
			"updated_at":  plan.UpdatedAt,
		}
	}

	namedQuery, args, err := sqlx.Named(queryCreatePaymentPlan, params)
	if err != nil {
		log.Printf("[CreatePaymentPlanInDB] sqlxNamed() got error: %v\nMetadata: %v\n", err, map[string]interface{}{"loan_id": loanID})
		return err
	}

	_, err = tx.ExecContext(ctxTimeout, r.db.Rebind(namedQuery), args...)
	if err != nil {
		log.Printf("[CreatePaymentPlanInDB] tx.ExecContext() got error: %v\nMetadata: %v\n", err, map[string]interface{}{"loan_id": loanID})
		return err
	}

	return nil
}
