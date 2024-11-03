package llm_implementations

import (
	"context"
	"fmt"
	"github.com/deployment-io/team-ai/enums/llm_implementation_enums"
	"github.com/deployment-io/team-ai/options/agent_options"
)

type BaseAgent struct {
	Role          string
	Backstory     string
	LLM           string
	MaxIterations int
}

//TODO Do will have a way to pass team execution options

type AgentInterface interface {
	Do(ctx context.Context, input string, opts ...agent_options.Execution) (string, error)
}

func Get(agentType llm_implementation_enums.LLMImplementationType, opts ...agent_options.Creation) (AgentInterface, error) {
	switch agentType {
	case llm_implementation_enums.OpenAIFunctionAgent:
		o, err := NewOpenAIFunctionAgent(opts...)
		if err != nil {
			return nil, err
		}
		return o, nil
	default:
		return nil, fmt.Errorf("agent type not supported")
	}
}
