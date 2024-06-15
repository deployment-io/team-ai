package agent_options

type Creation func(*CreationOption)

type CreationOption struct {
	Role          string
	Backstory     string
	Llm           string
	MaxIterations int
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

func WithMaxIterations(iterations int) Creation {
	return func(o *CreationOption) {
		o.MaxIterations = iterations
	}
}

func DefaultOpenAIFunction() *CreationOption {
	return &CreationOption{
		Llm: "gpt-4o",
	}
}
