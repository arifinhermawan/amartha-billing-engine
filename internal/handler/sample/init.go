package sample

import "fmt"

type sampleUseCaseProvider interface {
	HelloWorld()
}

type Handler struct {
	sample sampleUseCaseProvider
}

func NewHandler(sample sampleUseCaseProvider) *Handler {
	return &Handler{
		sample: sample,
	}
}

func (h *Handler) HelloWorld() {
	fmt.Println("HELLO FROM HANDLER")
	h.sample.HelloWorld()
}
