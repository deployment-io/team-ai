package agent_options

import (
	"github.com/ankit-arora/langchaingo/callbacks"
	"github.com/deployment-io/team-ai/rpc"
)

type Creation func(*CreationOption)

type CreationOption struct {
	Role            string
	Backstory       string
	Llm             string
	ApiVersion      string
	MaxIterations   int
	HttpClient      rpc.HTTPClientInterface
	CallbackHandler callbacks.Handler
}

func WithCallbackHandler(callbackHandler callbacks.Handler) Creation {
	return func(o *CreationOption) {
		o.CallbackHandler = callbackHandler
	}
}

func WithRole(role string) Creation {
	return func(o *CreationOption) {
		o.Role = role
	}
}

func WithBackstory(backstory string) Creation {
	return func(o *CreationOption) {
		o.Backstory = backstory
	}
}

func WithLLM(llm string) Creation {
	return func(o *CreationOption) {
		o.Llm = llm
	}
}

func WithApiVersion(apiVersion string) Creation {
	return func(o *CreationOption) {
		o.ApiVersion = apiVersion
	}
}

func WithMaxIterations(iterations int) Creation {
	return func(o *CreationOption) {
		o.MaxIterations = iterations
	}
}

func WithHttpClient(httpClient rpc.HTTPClientInterface) Creation {
	return func(o *CreationOption) {
		o.HttpClient = httpClient
	}
}

func DefaultOpenAIFunction() *CreationOption {
	return &CreationOption{
		Llm: "gpt-4o",
	}
}
