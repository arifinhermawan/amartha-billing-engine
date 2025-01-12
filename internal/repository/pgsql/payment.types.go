package pgsql

import "time"

// request
type UpdatePaymentPlanReq struct {
	PaymentID int64     `db:"payment_id"`
	IsPaid    bool      `db:"is_paid"`
	PaidDate  time.Time `db:"paid_date"`
}

// response
type PaymentPlan struct {
	ID         int64     `db:"id"`
	LoanID     int64     `db:"loan_id"`
	WeekNumber int       `db:"week_number"`
	Amount     float64   `db:"amount"`
	DueDate    time.Time `db:"due_date"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
