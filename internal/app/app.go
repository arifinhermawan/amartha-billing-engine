package app

import (
	"log"

	"github.com/arifinhermawan/amartha-billing-engine/internal/app/router"
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
	db, err := utils.InitDBConn(lib.GetConfig().Database)
	if err != nil {
		log.Fatalf("[NewApplication] utils.InitDBConn() got error: %v\n", err)
		return
	}

	// init app stack
	repo := server.NewRepositories(lib, db)

	// service
	svc := server.NewService(lib, repo)

	// usecase
	uc := server.NewUseCases(svc)

	// handler
	handlers := server.NewHandler(uc)
	router.HandleRequest(handlers)
}
