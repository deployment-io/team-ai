package agent_options

import (
	"github.com/ankit-arora/langchaingo/callbacks"
	"github.com/ankit-arora/langchaingo/prompts"
	"github.com/ankit-arora/langchaingo/schema"
	"github.com/ankit-arora/langchaingo/tools"
)

type Execution func(*ExecutionOption)

type ExecutionOption struct {
	Memory        schema.Memory
	ToolChoice    any
	Callback      callbacks.Handler
	Tools         []tools.Tool
	ExtraMessages []prompts.MessageFormatter
	JSONMode      bool
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

func WithJSONMode(jsonMode bool) Execution {
	return func(e *ExecutionOption) {
		e.JSONMode = jsonMode
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

func WithExtraMessages(messages []prompts.MessageFormatter) Execution {
	return func(e *ExecutionOption) {
		e.ExtraMessages = messages
	}
}
