package payment

import "github.com/go-playground/validator"

type payWeeklyInstallmentReq struct {
	LoanID int64 `json:"loan_id" validate:"required"`
}

func validate(r interface{}) error {
	validate := validator.New()
	return validate.Struct(r)
}
