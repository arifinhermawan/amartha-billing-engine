package server

import "github.com/arifinhermawan/project-template/internal/service/sample"

type Services struct {
	Sample *sample.Service
}

func NewService() *Services {
	return &Services{
		Sample: sample.NewService(),
	}
}
