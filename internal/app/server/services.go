package server

import (
	"github.com/arifinhermawan/amartha-billing-engine/internal/lib"
	"github.com/arifinhermawan/amartha-billing-engine/internal/service/auth"
	"github.com/arifinhermawan/amartha-billing-engine/internal/service/user"
)

type Services struct {
	Auth *auth.Service
	User *user.Service
}

func NewService(lib *lib.Lib, repo *Repositories) *Services {
	return &Services{
		Auth: auth.NewService(lib),
		User: user.NewService(lib, repo.DB),
	}
}
