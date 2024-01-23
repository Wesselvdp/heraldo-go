package main

import (
	"context"
	"fmt"

	"net/http"
	"os"

	// "heraldo-server/middleware"

	"github.com/gin-gonic/gin"

	openai "github.com/sashabaranov/go-openai"
)

type LLMInput struct {
	Prompt string `json:"prompt" binding:"required"`
}
type OaiResp struct {
	Answer string
	Usage  float64
}

func handleGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "poing",
	})
}

func handleLLM(c *gin.Context) {
	// Validate input
	var input LLMInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// singleStr := ""
	// scanner := bufio.NewScanner(strings.NewReader(input.Prompt))
	// for scanner.Scan() {
	// 	// fmt.Println(scanner.Text())
	// 	singleStr += scanner.Text()
	// }

	// fmt.Println("singleStr: " + singleStr)

	resp, err := callOpenAI(input.Prompt)

	// fmt.Println("Answer: " + answer)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"answer": resp.Answer,
		"usage":  resp.Usage,
	})
}

func callOpenAI(prompt string) (*OaiResp, error) {
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

func main() {
	// whiteList := map[string]bool{
	// 	"https://www.google.com": true,
	// 	"https://www.yahoo.com":  true,
	// }
	router := gin.Default()

	// Add whitelist middleware
	// router.Use(middleware.IPWhiteList(whiteList))
	router.GET("/albums", handleGet)

	router.POST("/llm", handleLLM)
	router.Run("0.0.0.0:8080")
}
