package agents

import (
	"fmt"
	"github.com/ankit-arora/langchaingo/tools"
	"github.com/deployment-io/team-ai/agents/agent_enums"
	"github.com/deployment-io/team-ai/options/agent_options"
)

type BaseAgent struct {
	Role          string
	Backstory     string
	LLM           string
	Tools         []tools.Tool
	MaxIterations int
}

//TODO Do will have a way to pass team execution options

type AgentInterface interface {
	Do(input string, opts ...agent_options.Execution) (string, error)
}

func Get(agentType agent_enums.AgentType, opts ...agent_options.Creation) (AgentInterface, error) {
	switch agentType {
	case agent_enums.OpenAIFunctionAgent:
		o, err := NewOpenAIFunctionAgent(opts...)
		if err != nil {
			return nil, err
		}
		return o, nil
	default:
		return nil, fmt.Errorf("agent type not supported")
	}
}
