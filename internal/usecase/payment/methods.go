package payment

import (
	"context"
	"errors"
	"log"
)

func (uc *UseCase) PayWeeklyInstallment(ctx context.Context, loanID int64) error {
	metadata := map[string]interface{}{
		"loan_id": loanID,
	}

	loan, err := uc.loan.GetLoanByID(ctx, loanID)
	if err != nil {
		log.Printf("[PayWeeklyInstallment] uc.loan.GetLoanByID() got error: %v\nMetadata: %v\n", err, metadata)
		return err
	}

	if loan.OutstandingBalance == 0 || !loan.IsActive {
		return errors.New("loan is inactive")
	}

	paymentAmount, err := uc.payment.PayWeeklyInstallment(ctx, loanID)
	if err != nil {
		log.Printf("[PayWeeklyInstallment] uc.payment.PayWeeklyInstallment() got error: %v\nMetadata: %v\n", err, metadata)
		return err
	}

	metadata["payment_amount"] = paymentAmount

	err = uc.loan.UpdateLoan(ctx, loanID, loan.OutstandingBalance-paymentAmount)
	if err != nil {
		log.Printf("[PayWeeklyInstallment] uc.loan.UpdateLoan() got error: %v\nMetadata: %v\n", err, metadata)
		return err
	}

	return nil
}
