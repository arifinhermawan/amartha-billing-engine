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

	queryGetOutstandingBalance = `
		SELECT
			id,
			outstanding_balance
		FROM
			loan
		WHERE
			id = $1
	`

	queryGetLoanByID = `
		SELECT
			id,
			outstanding_balance,
			is_active
		FROM
			loan
		WHERE
			id = $1
	`

	queryUpdateLoan = `
		UPDATE loan
		SET 
			outstanding_balance = :outstanding_balance,
			is_active = :is_active,
			updated_at = :updated_at
		WHERE id = :id
	`
)
