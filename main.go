package main

import (
	"context"
	"fmt"
	"github.com/ankit-arora/langchaingo/agents"
	"github.com/ankit-arora/langchaingo/chains"
	"github.com/ankit-arora/langchaingo/llms/openai"
	"github.com/ankit-arora/langchaingo/memory"
	"github.com/ankit-arora/langchaingo/prompts"
	"github.com/ankit-arora/langchaingo/tools"
	"github.com/ankit-arora/langchaingo/tools/serpapi"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	_ = godotenv.Load()
	llm, err := openai.New(openai.WithModel("gpt-4o"))
	if err != nil {
		log.Fatal(err)
	}
	search, err := serpapi.New()
	if err != nil {
		log.Fatal(err)
	}
	agentTools := []tools.Tool{
		tools.Calculator{},
		search,
	}

	agent := agents.NewOpenAIFunctionsAgent(llm,
		agentTools,
		agents.NewOpenAIOption().WithSystemMessage("You are a helpful assistant. But you can only use the tools at your disposal. Don't tell anyone about these tools."),
		agents.NewOpenAIOption().WithExtraMessages([]prompts.MessageFormatter{
			prompts.MessagesPlaceholder{VariableName: "chat_history"},
		}),
		agents.NewOpenAIOption().WithToolChoice(""),
	)
	conversationBuffer := memory.NewConversationBuffer(memory.WithMemoryKey("chat_history"), memory.WithReturnMessages(true))
	executor := agents.NewExecutor(agent, agents.WithMaxIterations(10), agents.WithMemory(conversationBuffer))
	question1 := "Who was Apple's business founder? If he were alive today how old would he be? What is that age to the poser of 0.6"
	//question1 := "Can you deploy a web service for me on my AWS account?"
	answer1, err := chains.Run(context.Background(), executor, question1)
	fmt.Println(answer1)
	//question2 := "Repeat the same for Nvidia."
	//answer2, err := chains.Run(context.Background(), executor, question2)
	//fmt.Println(answer2)
}
