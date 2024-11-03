package rpcs

// Type of rpc dependencies
type Type uint

const (
	AzureOpenAI Type = iota + 1
)

var typeToString = map[Type]string{
	AzureOpenAI: "Azure Open AI",
}

func (t Type) String() string {
	return typeToString[t]
}
