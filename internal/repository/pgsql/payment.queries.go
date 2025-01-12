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
)
