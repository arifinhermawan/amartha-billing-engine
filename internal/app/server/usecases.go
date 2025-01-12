package server

import "github.com/arifinhermawan/project-template/internal/usecase/sample"

type UseCases struct {
	Sample *sample.UseCase
}

func NewUseCases(svc *Services) *UseCases {
	return &UseCases{
		Sample: sample.NewUseCase(svc.Sample),
	}
}
