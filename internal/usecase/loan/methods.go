package loan

import (
	"context"
	"log"

	"github.com/arifinhermawan/amartha-billing-engine/internal/service/loan"
)

func (uc *UseCase) CreateLoan(ctx context.Context, req CreateLoanReq) error {
	metadata := map[string]interface{}{
		"user_id":           req.UserID,
		"amount":            req.Amount,
		"duration_in_weeks": req.DurationInWeeks,
		"interest_rate":     req.InterestRate,
	}

	_, err := uc.user.GetUserByID(ctx, req.UserID)
	if err != nil {
		log.Printf("[CreateLoan] uc.user.GetUserByID() got error: %v\nMetadata: %v\n", err, metadata)
		return err
	}

	err = uc.loan.CreateLoan(ctx, loan.CreateLoanReq(req))
	if err != nil {
		log.Printf("[CreateLoan] uc.loan.CreateLoan() got error: %v\nMetadata: %v\n", err, metadata)
		return err
	}

	return nil
}

func (uc *UseCase) GetOutstandingBalance(ctx context.Context, loanID int64) (float64, error) {
	outstandingBalance, err := uc.loan.GetOutstandingBalance(ctx, loanID)
	if err != nil {
		log.Printf("[GetOutstandingBalance] uc.loan.GetOutstandingBalance() got error: %v\nMetadata: %v\n", err, map[string]interface{}{"loan_id": loanID})
		return 0, err
	}

	return outstandingBalance, nil
}
