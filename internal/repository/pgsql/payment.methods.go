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

func (r *Repository) GetUnpaidPaymentPlanByLoanID(ctx context.Context, loanID int64) ([]PaymentPlan, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.lib.GetConfig().Database.DefaultTimeout)*time.Second)
	defer cancel()

	var plan []PaymentPlan
	err := r.db.SelectContext(ctxTimeout, &plan, queryGetUnpaidPaymentPlan, loanID)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("[GetUnpaidPaymentPlanByLoanID] r.db.SelectContext() got error: %v\nMetadata: %v\n", err, map[string]interface{}{"loan_id": loanID})
		return nil, err
	}

	return plan, nil
}

func (r *Repository) UpdatePaymentPlan(ctx context.Context, req UpdatePaymentPlanReq) error {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.lib.GetConfig().Database.DefaultTimeout)*time.Second)
	defer cancel()

	metadata := map[string]interface{}{
		"payment_id": req.PaymentID,
	}

	namedQuery, args, err := sqlx.Named(queryUpdatePaymentPlan, req)
	if err != nil {
		log.Printf("[UpdatePaymentPlan] sqlxNamed() got error: %v\nMetadata: %v\n", err, metadata)
		return err
	}

	_, err = r.db.ExecContext(ctxTimeout, r.db.Rebind(namedQuery), args...)
	if err != nil {
		log.Printf("[UpdatePaymentPlan] r.db.ExecContext() got error: %v\nMetadata: %v\n", err, metadata)
		return err
	}

	return nil

}
