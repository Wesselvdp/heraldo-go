package main

import (
	"fmt"
	"heraldo-server/pkg/handlers"
	"heraldo-server/pkg/middleware"

	// "heraldo-server/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// whiteList := map[string]bool{
	// 	"https://chamonix.netlify.app": true,
	// 	"http":  true,
	// }
	router := gin.Default()

	// Add whitelist middleware
	// router.Use(middleware.IPWhiteList(whiteList))
	router.Use(middleware.CORSMiddleware())

	router.GET("/health", handlers.Health)

	address := "127.0.0.1:8080"

	router.POST("/llm", handlers.ChatCompletion)
	fmt.Println("Server running on http://" + address + "...")
	router.Run(address)

}
