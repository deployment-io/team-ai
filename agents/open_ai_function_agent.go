package agents

import (
	"context"
	"github.com/ankit-arora/langchaingo/agents"
	"github.com/ankit-arora/langchaingo/chains"
	"github.com/ankit-arora/langchaingo/llms/openai"
)

type OpenAIFunctionAgent struct {
	BaseAgent
	llm *openai.LLM
}

func NewOpenAIFunctionAgent(opts ...CreationOption) (*OpenAIFunctionAgent, error) {
	options := defaultOpenAIFunctionAgentOptions()
	for _, opt := range opts {
		opt(options)
	}
	o := &OpenAIFunctionAgent{}
	o.Backstory = options.backstory
	o.LLM = options.llm
	o.Role = options.role
	o.Goal = options.goal
	o.Tools = options.tools
	o.MaxIterations = options.maxIterations
	var err error
	o.llm, err = openai.New(openai.WithModel(o.LLM))
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *OpenAIFunctionAgent) Do() (string, error) {
	//ignore memory for now
	a := agents.NewOpenAIFunctionsAgent(o.llm,
		o.Tools,
		agents.NewOpenAIOption().WithSystemMessage(o.Backstory),
		//agents.NewOpenAIOption().WithExtraMessages([]prompts.MessageFormatter{
		//	prompts.MessagesPlaceholder{VariableName: "chat_history"},
		//}),
		agents.NewOpenAIOption().WithToolChoice(""),
	)
	//agents.WithMemory(conversationBuffer)
	executor := agents.NewExecutor(a, agents.WithMaxIterations(o.MaxIterations))
	answer, err := chains.Run(context.Background(), executor, o.Goal)
	return answer, err
}
