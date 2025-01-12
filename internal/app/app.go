package app

import (
	"github.com/arifinhermawan/project-template/internal/app/server"
	"github.com/arifinhermawan/project-template/internal/app/utils"
	"github.com/arifinhermawan/project-template/internal/lib"
	"github.com/arifinhermawan/project-template/internal/lib/configuration"
	"github.com/arifinhermawan/project-template/internal/lib/time"
)

func NewApplication() {
	// initialize lib
	lib := lib.New(
		configuration.New(),
		time.New(),
	)

	// init db connection
	_, _ = utils.InitDBConn(lib.GetConfig().Database)

	// init redis connection
	_, _ = utils.InitRedisConn(lib.GetConfig().Redis)

	// init app stack
	// service
	svc := server.NewService()

	// usecase
	uc := server.NewUseCases(svc)

	// handler
	handler := server.NewHandler(uc)

	// test run
	handler.Sample.HelloWorld()
}
