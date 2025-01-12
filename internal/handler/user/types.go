package user

import "github.com/go-playground/validator"

// Request
type createUserReq struct {
	Name     string `json:"name" validate:"required,min=3,max=30"`
	Password string `json:"password" validate:"required"`
}

func validate(r interface{}) error {
	validate := validator.New()
	return validate.Struct(r)
}

// Response
type delinquentUser struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
