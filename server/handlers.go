package server

import (
	"bytes"
	"io"
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

		log.Println("Headers:", c.Request.Header)

		bodyBytes, _ := io.ReadAll(c.Request.Body)
		log.Println("Raw Body:", string(bodyBytes))

		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		if err := c.ShouldBindJSON(&payload); err != nil {
			log.Println("JSON bind error:", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  "Invalid JSON payload",
				"detail": err.Error(),
			})
			return
		}

		log.Println("Parsed payload:", payload)

		c.JSON(http.StatusCreated, gin.H{
			"message": "Payload received successfully",
			"payload": payload,
		})
	}
}
