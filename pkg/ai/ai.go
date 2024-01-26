package ai

import (
	"context"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

type OaiResp struct {
	Answer string
	Usage  float64
}

func CallOpenAI(prompt string) (*OaiResp, error) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       openai.GPT4,
			Temperature: 0.4,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	response := &OaiResp{}

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return nil, err
	}

	response.Answer = resp.Choices[0].Message.Content
	response.Usage = float64(resp.Usage.TotalTokens)

	return response, nil
}
