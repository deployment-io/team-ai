package agent_enums

type AgentType uint

const (
	DevOpsUserAgent AgentType = iota + 1
	ServiceUserAgent
	AutomationAgent
	RouterAgent  //routes a message to an appropriate agent in the app server
	GenericAgent //runs the code of the selected agent in the app server
	AgentRunner
	MaxAgentType //always add new types before MaxAgentType
)

var typeToString = map[AgentType]string{
	DevOpsUserAgent:  "DevOps User Agent",
	ServiceUserAgent: "Service User Agent",
	AutomationAgent:  "Automation Agent",
	RouterAgent:      "Router Agent",
	GenericAgent:     "Generic Agent",
	AgentRunner:      "Agent Runner",
}

func (t AgentType) String() string {
	return typeToString[t]
}
