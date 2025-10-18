package services

import (
	"context"
	"log"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func initAi() *openai.Client {
	key := os.Getenv("OPENAI_API_KEY")

	client := openai.NewClient(key)
	return client
}

func AskAI(prompt string) string {
	client := initAi()
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: "gpt-5",
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			N: 1, // Number of request choices
		},
	)

	if err != nil {
		log.Printf("OpenAI error: %v", err)
		return "Failed to get response from AI."
	}

	if len(resp.Choices) == 0 {
		log.Println("OpenAI returned no choices.")
		return "No response received."
	}

	return resp.Choices[0].Message.Content
}

func AskAIWithContext(prompt string) string {
	client := initAi()

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4o, // or GPT4
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a senior Go engineer who gives concise, accurate answers.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			N: 1, // Number of request choices
		},
	)

	if err != nil {
		log.Printf("OpenAI error: %v", err)
		return "Failed to get response from AI."
	}

	if len(resp.Choices) == 0 {
		log.Println("OpenAI returned no choices.")
		return "No response received."
	}

	return resp.Choices[0].Message.Content
}
