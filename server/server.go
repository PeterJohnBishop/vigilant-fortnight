package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	AddWebhookRoutes(router)
	log.Println("Server listening on :8080")
	router.Run(":8080")
}
