package agent_enums

type AgentType uint

const (
	DevOpsUserAgent AgentType = iota + 1
	ServiceUserAgent
	MaxType //always add new types before MaxType
)

var typeToString = map[AgentType]string{
	DevOpsUserAgent:  "DevOps User Agent",
	ServiceUserAgent: "Service User Agent",
}

func (t AgentType) String() string {
	return typeToString[t]
}
