package lib

import (
	"time"

	"github.com/arifinhermawan/amartha-billing-engine/internal/lib/configuration"
)

type configProvider interface {
	// GetConfig retrieves the application configuration in a thread-safe manner.
	// It applies singleton pattern so it ensures that the configuration is loaded only once,
	// even if the function is called multiple times. The configuration is loaded using the loadConfig()
	// function, which loads it from a YAML file based on the environment setting.
	GetConfig() *configuration.AppConfig
}

type timeProvider interface {
	// GetTimeGMT7 retrieves the current time in the GMT+7.
	GetTimeGMT7() time.Time
}

type Lib struct {
	config configProvider
	time   timeProvider
}

func New(config configProvider, time timeProvider) *Lib {
	return &Lib{
		config: config,
		time:   time,
	}
}
