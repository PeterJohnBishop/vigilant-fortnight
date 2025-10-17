package server

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	AddBasicRoutes(router)
	AddWebhookRoutes(router)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port if PORT is not set
	}
	log.Printf("Server listening on :%s\n", port)
	err := router.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
