package router_agent

import (
	"os"

	"github.com/ankit-arora/langchaingo/callbacks"
	"github.com/deployment-io/team-ai/enums/llm_implementation_enums"
	"github.com/deployment-io/team-ai/enums/rpcs"
	"github.com/deployment-io/team-ai/llm_implementations"
	"github.com/deployment-io/team-ai/options/agent_options"
	"github.com/deployment-io/team-ai/rpc"
)

const role = "router agent"

const backstory = `You're a friendly AI assistant who can only use the tools provided. Don't assume anything. Only use the tools.
If you're asked to do something else or asked any other question, only give a summary of what you can do for the user today.
Send all responses in Markdown format.`

const maxIterations = 10

func New(llm, llmApiVersion, extraContext string, callbackHandler callbacks.Handler) (llm_implementations.AgentInterface, error) {
	agentBackstory := os.Getenv("ROUTER_AGENT_BACKSTORY")
	if len(agentBackstory) == 0 {
		agentBackstory = backstory
	}
	if len(extraContext) > 0 {
		agentBackstory += "\n\n" + extraContext
	}
	httpClient := rpc.NewHTTPClient(rpcs.AzureOpenAI, true, true, 2)
	return llm_implementations.Get(llm_implementation_enums.OpenAIFunctionAgent, agent_options.WithBackstory(agentBackstory),
		agent_options.WithRole(role),
		agent_options.WithMaxIterations(maxIterations),
		agent_options.WithLLM(llm),
		agent_options.WithApiVersion(llmApiVersion),
		agent_options.WithHttpClient(httpClient),
		agent_options.WithCallbackHandler(callbackHandler),
	)
}
