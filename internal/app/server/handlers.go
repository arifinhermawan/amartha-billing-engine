package server

import "github.com/arifinhermawan/amartha-billing-engine/internal/handler/user"

type Handlers struct {
	User *user.Handler
}

func NewHandler(uc *UseCases) *Handlers {
	return &Handlers{
		User: user.NewHandler(uc.User),
	}
}
