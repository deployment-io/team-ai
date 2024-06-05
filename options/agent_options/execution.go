package agent_options

import (
	"github.com/ankit-arora/langchaingo/callbacks"
	"github.com/ankit-arora/langchaingo/schema"
)

type Execution func(*ExecutionOption)

type ExecutionOption struct {
	Memory     schema.Memory
	ToolChoice any
	Callback   callbacks.Handler
}

func WithMemory(m schema.Memory) Execution {
	return func(e *ExecutionOption) {
		e.Memory = m
	}
}

func WithToolChoice(t any) Execution {
	return func(e *ExecutionOption) {
		e.ToolChoice = t
	}
}

func WithCallback(callback callbacks.Handler) Execution {
	return func(e *ExecutionOption) {
		e.Callback = callback
	}
}
