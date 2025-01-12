package loan

import (
	"context"
	"log"
	"time"

	"github.com/arifinhermawan/amartha-billing-engine/internal/repository/pgsql"
)

func (svc *Service) CreateLoan(ctx context.Context, req CreateLoanReq) error {
	metadata := map[string]interface{}{
		"user_id":           req.UserID,
		"amount":            req.Amount,
		"duration_in_weeks": req.DurationInWeeks,
		"interest_rate":     req.InterestRate,
	}

	outstandingBalance := req.Amount + (req.Amount * req.InterestRate)
	metadata["outstanding_balance"] = outstandingBalance

	tx, err := svc.db.BeginTX(ctx, nil)
	if err != nil {
		log.Printf("[CreateLoan] svc.db.BeginTX() got error: %v\nMetadata: %v\n", err, metadata)
		return err
	}

	defer func() {
		if err != nil {
			errRollback := tx.Rollback()
			if errRollback != nil {
				log.Printf("[CreateLoan] tx.Rollback() got error: %v\nMetadata: %v\n", err, metadata)
			}
		}
	}()

	timeNow := svc.lib.GetTimeGMT7()
	eod := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 23, 59, 59, 59, timeNow.Location())
	loanID, err := svc.db.CreateLoanInDB(ctx, tx, pgsql.CreateLoanReq{
		UserID:             req.UserID,
		Amount:             req.Amount,
		OutstandingBalance: outstandingBalance,
		InterestRate:       req.InterestRate,
		StartDate:          timeNow,
		EndDate:            eod.Add(time.Duration(req.DurationInWeeks) * 7 * 24 * time.Hour),
		CreatedAt:          timeNow,
		UpdatedAt:          timeNow,
	})
	if err != nil {
		log.Printf("[CreateLoan] svc.db.CreateLoanInDB() got error: %v\nMetadata: %v\n", err, metadata)
		return err
	}

	metadata["loan_id"] = loanID

	weeklyPayment := outstandingBalance / float64(req.DurationInWeeks)
	metadata["weekly_payment"] = weeklyPayment

	paymentPlans := make([]pgsql.PaymentPlan, req.DurationInWeeks)
	for i := 0; i < req.DurationInWeeks; i++ {
		dueDate := timeNow.Add(time.Duration(i+1) * 7 * 24 * time.Hour)
		dueDateEOD := time.Date(dueDate.Year(), dueDate.Month(), dueDate.Day(), 23, 59, 59, 59, timeNow.Location())

		paymentPlans[i] = pgsql.PaymentPlan{
			LoanID:     loanID,
			WeekNumber: i + 1,
			Amount:     weeklyPayment,
			DueDate:    dueDateEOD,
			CreatedAt:  timeNow,
			UpdatedAt:  timeNow,
		}
	}

	err = svc.db.CreatePaymentPlanInDB(ctx, tx, paymentPlans)
	if err != nil {
		log.Printf("[CreateLoan] svc.db.CreatePaymentPlanInDB() got error: %v\nMetadata: %v\n", err, metadata)
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("[CreateLoan] tx.Commit() got error: %v\nMetadata: %v\n", err, metadata)
		return err
	}

	return nil
}
