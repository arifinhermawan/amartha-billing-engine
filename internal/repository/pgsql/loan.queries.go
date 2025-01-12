package pgsql

const (
	queryCreateLoan = `
		INSERT INTO loan(user_id, amount, interest_rate, outstanding_balance, start_date, end_date, created_at, updated_at)
		VALUES(
			:user_id,
			:amount,
			:interest_rate,
			:outstanding_balance,
			:start_date,
			:end_date,
			:created_at,
			:updated_at
		)
		RETURNING id
	`
)
