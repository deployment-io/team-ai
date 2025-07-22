package agents

import (
	"context"
	"fmt"
	"github.com/ankit-arora/langchaingo/callbacks"
	"github.com/ankit-arora/langchaingo/schema"
	"github.com/ankit-arora/langchaingo/tools"
	"github.com/deployment-io/team-ai/agents/automation_agent"
	"github.com/deployment-io/team-ai/agents/devops_user_agent"
	"github.com/deployment-io/team-ai/agents/service_user_agent"
	"github.com/deployment-io/team-ai/enums/agent_enums"
	"github.com/deployment-io/team-ai/llm_implementations"
	"github.com/deployment-io/team-ai/options/agent_options"
	"log"
)

const GroupIDContextKey = "groupID"
const ThreadIDContextKey = "threadID"
const TokenContextKey = "token"
const OrganizationIDContextKey = "organizationID"
const InsertedAssistantMessageIDContextKey = "assistantMessageID"
const DeploymentIDContextKey = "deploymentID"
const InputMessageIDKey = "inputMessageID"

func GetAgentToAssist(agentType agent_enums.AgentType, llm, llmApiVersion, extraContext string, callbackHandler callbacks.Handler) (llm_implementations.AgentInterface, error) {
	switch agentType {
	case agent_enums.DevOpsUserAgent:
		devopsAgent, err := devops_user_agent.New(llm, llmApiVersion, extraContext)
		if err != nil {
			return nil, err
		}
		return devopsAgent, nil
	case agent_enums.ServiceUserAgent:
		serviceAgent, err := service_user_agent.New(llm, llmApiVersion, extraContext)
		if err != nil {
			return nil, err
		}
		return serviceAgent, nil
	case agent_enums.AutomationAgent:
		automationAgent, err := automation_agent.New(llm, llmApiVersion, extraContext, callbackHandler)
		if err != nil {
			return nil, err
		}
		return automationAgent, nil
	default:
		return nil, fmt.Errorf("agent type %s not supported", agentType)
	}
}

func Assist(ctx context.Context, input string, memory schema.Memory, callback callbacks.Handler,
	agentType agent_enums.AgentType,
	agent llm_implementations.AgentInterface, assistantTools map[agent_enums.AgentType][]tools.Tool) (string, error) {
	log.Printf("Assistant assist called with input %s", input)
	out, err := agent.Do(ctx, input, agent_options.WithMemory(memory), agent_options.WithToolChoice("auto"),
		agent_options.WithCallback(callback), agent_options.WithTools(assistantTools[agentType]))
	if err != nil {
		return "", err
	}
	return out, nil
}
