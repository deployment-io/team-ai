package agents

import (
	"context"
	"github.com/ankit-arora/langchaingo/agents"
	"github.com/ankit-arora/langchaingo/chains"
	"github.com/ankit-arora/langchaingo/llms"
	"github.com/ankit-arora/langchaingo/llms/openai"
	"github.com/ankit-arora/langchaingo/prompts"
	"github.com/deployment-io/team-ai/options/agent_options"
)

type OpenAIFunctionAgent struct {
	BaseAgent
	llm llms.Model
}

func NewOpenAIFunctionAgent(opts ...agent_options.Creation) (*OpenAIFunctionAgent, error) {
	options := agent_options.DefaultOpenAIFunction()
	for _, opt := range opts {
		opt(options)
	}
	o := &OpenAIFunctionAgent{}
	o.Backstory = options.Backstory
	o.LLM = options.Llm
	o.Role = options.Role
	o.Tools = options.Tools
	o.MaxIterations = options.MaxIterations
	var err error
	o.llm, err = openai.New(openai.WithModel(o.LLM))
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *OpenAIFunctionAgent) Do(input string, opts ...agent_options.Execution) (string, error) {
	options := &agent_options.ExecutionOption{}
	for _, opt := range opts {
		opt(options)
	}
	var agentOptions = []agents.Option{
		agents.NewOpenAIOption().WithSystemMessage(o.Backstory),
		agents.NewOpenAIOption().WithToolChoice(options.ToolChoice),
		agents.WithCallbacksHandler(options.Callback),
	}
	if options.Memory != nil {
		agentOptions = append(agentOptions, agents.NewOpenAIOption().WithExtraMessages([]prompts.MessageFormatter{
			prompts.MessagesPlaceholder{VariableName: "chat_history"},
		}))
	}

	a := agents.NewOpenAIFunctionsAgent(o.llm,
		o.Tools,
		agentOptions...,
	)

	var executionOptions = []agents.Option{
		agents.WithMaxIterations(o.MaxIterations),
	}

	if options.Memory != nil {
		executionOptions = append(agentOptions, agents.WithMemory(options.Memory))
	}

	executor := agents.NewExecutor(a, executionOptions...)
	answer, err := chains.Run(context.Background(), executor, input)
	return answer, err
}
