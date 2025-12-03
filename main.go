package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/sashabaranov/go-openai"
)

func main() {
	// 環境変数からAPIキーを取得
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("OPENAI_API_KEY is not set")
		os.Exit(1)
	}

	// OpenAIクライアントを初期化
	client := openai.NewClient(apiKey)

	fmt.Println("nebula - OpenAI Chat CLI")
	fmt.Println("Type 'exit' or 'quit' to end the conversation")
	fmt.Println("--------------------------------")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("You: ")
		if !scanner.Scan() {
			break
		}
		userInput := strings.TrimSpace(scanner.Text())

		// 終了コマンドをチェック
		if userInput == "exit" || userInput == "quit" {
			fmt.Println("Goodbye!")
			break
		}

		if userInput == "" {
			continue
		}

		// OpenAI API に送信
		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT4Dot1Nano,
				Messages: []openai.ChatCompletionMessage{
					{
						Role: openai.ChatMessageRoleUser,
						Content: userInput,
					},
				},
			},
		)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		if len(resp.Choices) > 0 {
			fmt.Printf("Assistant: %s\n\n", resp.Choices[0].Message.Content)
		} else {
			fmt.Println("No response received from OpenAI")
		}
	}
}