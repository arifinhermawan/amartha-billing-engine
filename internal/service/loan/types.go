package loan

// request
type CreateLoanReq struct {
	UserID          int64
	Amount          float64
	InterestRate    float64
	DurationInWeeks int
}

// response
type Loan struct {
	ID                 int64
	IsActive           bool
	OutstandingBalance float64
}
