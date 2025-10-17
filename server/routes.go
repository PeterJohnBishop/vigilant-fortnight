package server

import "github.com/gin-gonic/gin"

func AddBasicRoutes(r *gin.Engine) {
	r.GET("/", Greet())
	r.GET("/health", Health())
}

func AddWebhookRoutes(r *gin.Engine) {
	r.POST("/webhook", HandlePayload())
}
