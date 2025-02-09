package service_user_agent

import (
	"github.com/deployment-io/team-ai/enums/llm_implementation_enums"
	"github.com/deployment-io/team-ai/enums/rpcs"
	"github.com/deployment-io/team-ai/llm_implementations"
	"github.com/deployment-io/team-ai/options/agent_options"
	"github.com/deployment-io/team-ai/rpc"
	"os"
)

const role = "service user agent"

const backstory = `You're a friendly AI assistant who can only use the tools provided. Don't assume anything. Only use the tools.
If you're asked to do something else or asked any other question, only give a summary of what you can do for the user today.
Send all responses in Markdown format.`

const maxIterations = 10

func New(llm, extraContext string) (llm_implementations.AgentInterface, error) {
	serviceBackstory := os.Getenv("SERVICE_AGENT_BACKSTORY")
	if len(serviceBackstory) == 0 {
		serviceBackstory = backstory
	}
	if len(extraContext) > 0 {
		serviceBackstory += "\n" + "Here is some extra context:\n" + extraContext
	}
	httpClient := rpc.NewHTTPClient(rpcs.AzureOpenAI, true, true, 5)
	return llm_implementations.Get(llm_implementation_enums.OpenAIFunctionAgent, agent_options.WithBackstory(serviceBackstory),
		agent_options.WithRole(role),
		agent_options.WithMaxIterations(maxIterations),
		agent_options.WithLLM(llm),
		agent_options.WithHttpClient(httpClient))
}
