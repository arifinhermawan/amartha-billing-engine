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
