package handlers

import (
	"net/http"

	"heraldo-server/pkg/ai"

	"github.com/gin-gonic/gin"
)

type LLMInput struct {
	Prompt string `json:"prompt" binding:"required"`
}

func Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "poing",
	})
}

func ChatCompletion(c *gin.Context) {
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

	resp, err := ai.CallOpenAI(input.Prompt)

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
