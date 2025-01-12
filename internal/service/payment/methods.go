package payment

import (
	"context"
	"log"

	"github.com/arifinhermawan/amartha-billing-engine/internal/repository/pgsql"
)

func (svc *Service) PayWeeklyInstallment(ctx context.Context, loanID int64) (float64, error) {
	metadata := map[string]interface{}{
		"loan_id": loanID,
	}

	plans, err := svc.db.GetUnpaidPaymentPlanByLoanID(ctx, loanID)
	if err != nil {
		log.Printf("[PayWeeklyInstallment] svc.db.GetUnpaidPaymentPlanByLoanID() got error: %v\nMetadata: %v\n", err, metadata)
		return 0, err
	}

	if len(plans) == 0 {
		return 0, nil
	}

	plan := plans[0]

	metadata["payment_id"] = plan.ID
	metadata["payment_amount"] = plan.Amount

	err = svc.db.UpdatePaymentPlan(ctx, pgsql.UpdatePaymentPlanReq{
		PaymentID: plan.ID,
		IsPaid:    true,
		PaidDate:  svc.lib.GetTimeGMT7(),
	})
	if err != nil {
		log.Printf("[PayWeeklyInstallment] svc.db.UpdatePaymentPlan() got error: %v\nMetadata: %v\n", err, metadata)
		return 0, err
	}

	return plan.Amount, nil
}
