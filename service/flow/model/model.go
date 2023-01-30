package model

type FlowStatus int

const (
	Proceed FlowStatus = 1
	Cancel  FlowStatus = 2
)

type StatusStream struct {
	Forward chan FlowStatus
	Back    chan FlowStatus
}

func NewStatusStream() *StatusStream {
	stream := StatusStream{
		Forward: make(chan FlowStatus),
		Back:    make(chan FlowStatus),
	}
	return &stream
}
