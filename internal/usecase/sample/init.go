package sample

import "fmt"

type sampleServiceProvider interface {
	HelloWorld()
}

type UseCase struct {
	sample sampleServiceProvider
}

func NewUseCase(sample sampleServiceProvider) *UseCase {
	return &UseCase{
		sample: sample,
	}
}

func (u *UseCase) HelloWorld() {
	fmt.Println("HELLO FROM USECASE")
	u.sample.HelloWorld()
}
