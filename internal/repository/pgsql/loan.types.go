package pgsql

import "time"

// request
type CreateLoanReq struct {
	UserID             int64     `db:"user_id"`
	Amount             float64   `db:"amount"`
	OutstandingBalance float64   `db:"outstanding_balance"`
	InterestRate       float64   `db:"interest_rate"`
	StartDate          time.Time `db:"start_date"`
	EndDate            time.Time `db:"end_date"`
	CreatedAt          time.Time `db:"created_at"`
	UpdatedAt          time.Time `db:"updated_at"`
}

// response
type Loan struct {
	ID                 int64   `db:"id"`
	OutstandingBalance float64 `db:"outstanding_balance"`
}
