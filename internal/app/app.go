package app

import (
	"github.com/arifinhermawan/amartha-billing-engine/internal/app/server"
	"github.com/arifinhermawan/amartha-billing-engine/internal/app/utils"
	"github.com/arifinhermawan/amartha-billing-engine/internal/lib"
	"github.com/arifinhermawan/amartha-billing-engine/internal/lib/configuration"
	"github.com/arifinhermawan/amartha-billing-engine/internal/lib/time"
)

func NewApplication() {
	// initialize lib
	lib := lib.New(
		configuration.New(),
		time.New(),
	)

	// init db connection
	_, _ = utils.InitDBConn(lib.GetConfig().Database)

	// init app stack
	repo := server.NewRepositories()

	// service
	svc := server.NewService(repo)

	// usecase
	uc := server.NewUseCases(svc)

	// handler
	_ = server.NewHandler(uc)
}
