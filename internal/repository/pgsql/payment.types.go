package pgsql

import "time"

type PaymentPlan struct {
	LoanID     int64     `db:"loan_id"`
	WeekNumber int       `db:"week_number"`
	Amount     float64   `db:"amount"`
	DueDate    time.Time `db:"due_date"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
