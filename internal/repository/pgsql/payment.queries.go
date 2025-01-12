package pgsql

const (
	queryCreatePaymentPlan = `
		INSERT INTO payment(loan_id, week_number, amount, due_date, created_at, updated_at)
		VALUES(
			:loan_id,
			:week_number,
			:amount,
			:due_date,
			:created_at,
			:updated_at
		)
	`

	queryGetUnpaidPaymentPlan = `
		SELECT 
			id,
			amount
		FROM
			payment
		WHERE
			loan_id = $1
		AND
			is_paid = false
		ORDER BY 
			week_number ASC
	`

	queryUpdatePaymentPlan = `
		UPDATE payment
		SET
			is_paid = :is_paid,
			paid_date = :paid_date,
			updated_at = :paid_date
		WHERE
			id = :payment_id
	`
)
