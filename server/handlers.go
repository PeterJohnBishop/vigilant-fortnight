package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Greet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"message": "Hello from Vigilnt-Fortnight",
		})
	}
}

func Health() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"message": "Server is up.",
		})
	}
}

func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload map[string]interface{}

		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON payload",
			})
			return
		}

		log.Println(payload)

		c.JSON(http.StatusCreated, gin.H{
			"message": "Payload received successfully",
			"payload": payload,
		})
	}
}
