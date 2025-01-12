package loan

import "github.com/go-playground/validator"

type createLoanReq struct {
	UserID          int64   `json:"user_id" validate:"required"`
	Amount          float64 `json:"amount" validate:"required"`
	InterestRate    float64 `json:"interest_rate" validate:"required"`
	DurationInWeeks int     `json:"duration_in_weeks"`
}

func validate(r interface{}) error {
	validate := validator.New()
	return validate.Struct(r)
}
