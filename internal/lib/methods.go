package lib

import (
	"time"

	"github.com/arifinhermawan/project-template/internal/lib/configuration"
)

func (i *Lib) GetConfig() *configuration.AppConfig {
	return i.config.GetConfig()
}

func (i *Lib) GetTimeGMT7() time.Time {
	return i.time.GetTimeGMT7()
}