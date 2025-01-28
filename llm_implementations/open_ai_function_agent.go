package llm_implementations

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
	o.MaxIterations = options.MaxIterations
	var err error
	o.llm, err = openai.New(openai.WithModel(o.LLM), openai.WithAPIType(openai.APITypeAzure),
		openai.WithAPIVersion("2024-02-01"), openai.WithHTTPClient(options.HttpClient),
		openai.WithCallback(options.CallbackHandler))
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *OpenAIFunctionAgent) Do(ctx context.Context, input string, opts ...agent_options.Execution) (string, error) {
	options := &agent_options.ExecutionOption{}
	for _, opt := range opts {
		opt(options)
	}
	var agentOptions = []agents.Option{
		agents.NewOpenAIOption().WithSystemMessage(o.Backstory),
		agents.NewOpenAIOption().WithToolChoice(options.ToolChoice),
		agents.NewOpenAIOption().WithJSONMode(options.JSONMode),
		agents.WithCallbacksHandler(options.Callback),
	}
	if options.ExtraMessages != nil {
		extraMessages := options.ExtraMessages
		if options.Memory != nil {
			extraMessages = append(extraMessages, prompts.MessagesPlaceholder{VariableName: "chat_history"})
		}
		agentOptions = append(agentOptions, agents.NewOpenAIOption().WithExtraMessages(extraMessages))
	} else {
		if options.Memory != nil {
			agentOptions = append(agentOptions, agents.NewOpenAIOption().WithExtraMessages([]prompts.MessageFormatter{
				prompts.MessagesPlaceholder{VariableName: "chat_history"},
			}))
		}
	}

	a := agents.NewOpenAIFunctionsAgent(o.llm,
		options.Tools,
		agentOptions...,
	)

	var executionOptions = []agents.Option{
		agents.WithMaxIterations(o.MaxIterations),
	}

	if options.Memory != nil {
		executionOptions = append(agentOptions, agents.WithMemory(options.Memory))
	}

	if options.Callback != nil {
		executionOptions = append(executionOptions, agents.WithCallbacksHandler(options.Callback))
	}

	executor := agents.NewExecutor(a, executionOptions...)
	answer, err := chains.Run(ctx, executor, input)
	return answer, err
}
