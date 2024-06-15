package agent_options

import (
	"github.com/ankit-arora/langchaingo/callbacks"
	"github.com/ankit-arora/langchaingo/schema"
	"github.com/ankit-arora/langchaingo/tools"
)

type Execution func(*ExecutionOption)

type ExecutionOption struct {
	Memory     schema.Memory
	ToolChoice any
	Callback   callbacks.Handler
	Tools      []tools.Tool
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

func WithTools(tools []tools.Tool) Execution {
	return func(e *ExecutionOption) {
		e.Tools = tools
	}
}
