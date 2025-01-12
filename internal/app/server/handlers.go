package server

import "github.com/arifinhermawan/project-template/internal/handler/sample"

type Handlers struct {
	Sample *sample.Handler
}

func NewHandler(uc *UseCases) *Handlers {
	return &Handlers{
		Sample: sample.NewHandler(uc.Sample),
	}
}
