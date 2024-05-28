package agents

import (
	"fmt"
	"github.com/ankit-arora/langchaingo/tools"
	"github.com/deployment-io/team-ai/agents/agent_enums"
)

type BaseAgent struct {
	Role          string
	Goal          string
	Backstory     string
	LLM           string
	Tools         []tools.Tool
	MaxIterations int
}

//TODO Do will have a way to pass team execution options

type AgentInterface interface {
	Do() (string, error)
}

func Get(agentType agent_enums.AgentType, opts ...CreationOption) (AgentInterface, error) {
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

type CreationOption func(*options)

type options struct {
	role          string
	goal          string
	backstory     string
	llm           string
	maxIterations int
	tools         []tools.Tool
}

func WithRole(role string) CreationOption {
	return func(o *options) {
		o.role = role
	}
}

func WithGoal(goal string) CreationOption {
	return func(o *options) {
		o.goal = goal
	}
}

func WithBackstory(backstory string) CreationOption {
	return func(o *options) {
		o.backstory = backstory
	}
}

func WithLLM(llm string) CreationOption {
	return func(o *options) {
		o.llm = llm
	}
}

func WithMaxIterations(iterations int) CreationOption {
	return func(o *options) {
		o.maxIterations = iterations
	}
}

func WithTools(tools []tools.Tool) CreationOption {
	return func(o *options) {
		o.tools = tools
	}
}

func defaultOpenAIFunctionAgentOptions() *options {
	return &options{
		llm: "gpt-4o",
	}
}
