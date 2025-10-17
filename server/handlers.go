package server

import (
	"bytes"
	"encoding/json"
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

func HandlePayload() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload map[string]interface{}

		// Log headers as pretty JSON
		logAsPrettyJSON("Headers", c.Request.Header)

		// Read and log raw body
		bodyBytes, _ := io.ReadAll(c.Request.Body)
		log.Printf("Raw Body:\n%s", string(bodyBytes))

		// Reset body for further reading
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// Bind and log parsed JSON payload
		if err := c.ShouldBindJSON(&payload); err != nil {
			log.Printf("JSON bind error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  "Invalid JSON payload",
				"detail": err.Error(),
			})
			return
		}

		logAsPrettyJSON("Parsed Payload", payload)

		// Respond
		c.JSON(http.StatusCreated, gin.H{
			"message": "Payload received successfully",
			"payload": payload,
		})
	}
}

func logAsPrettyJSON(label string, data interface{}) {
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Printf("%s: error marshaling JSON: %v", label, err)
		return
	}
	log.Printf("%s:\n%s", label, string(jsonBytes))
}
