package main

import (
	"fmt"
	"github.com/ankit-arora/langchaingo/tools"
	"github.com/ankit-arora/langchaingo/tools/serpapi"
	"github.com/deployment-io/team-ai/agents"
	"github.com/deployment-io/team-ai/agents/agent_enums"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	_ = godotenv.Load()

	backstory := `You're an AI researcher at a large company.
	You're responsible for finding the best innovative company in AI and tell us why.
Reply with only one company with the reason why you think it's most innovative. Use the tools provided if you need help.`
	search, err := serpapi.New()
	if err != nil {
		log.Fatal(err)
	}
	agentTools := []tools.Tool{
		search,
		tools.Calculator{},
	}
	a, err := agents.Get(agent_enums.OpenAIFunctionAgent, agents.WithBackstory(backstory),
		agents.WithMaxIterations(10), agents.WithRole("AI researcher"),
		agents.WithGoal("Find the best AI company. Get its valuation and divide it by 5"), agents.WithTools(agentTools))
	if err != nil {
		panic(err)
	}
	out, err := a.Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)

}