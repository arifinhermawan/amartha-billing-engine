package sample

import "fmt"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) HelloWorld() {
	fmt.Println("HELLO FROM SERVICE")
}
