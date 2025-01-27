package agent_enums

type AgentType uint

const (
	DevOpsUserAgent AgentType = iota + 1
	ServiceUserAgent
	AutomationAgent
	MaxAgentType //always add new types before MaxAgentType
)

var typeToString = map[AgentType]string{
	DevOpsUserAgent:  "DevOps User Agent",
	ServiceUserAgent: "Service User Agent",
	AutomationAgent:  "Automation Agent",
}

func (t AgentType) String() string {
	return typeToString[t]
}
