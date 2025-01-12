package auth

import (
	"time"

	"github.com/arifinhermawan/amartha-billing-engine/internal/lib/configuration"
)

type libProvider interface {
	GetConfig() *configuration.AppConfig
	GetTimeGMT7() time.Time
}

type Service struct {
	lib libProvider
}

func NewService(lib libProvider) *Service {
	return &Service{
		lib: lib,
	}
}
