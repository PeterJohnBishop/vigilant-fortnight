package server

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"vigilant-fortnight/models"

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

func logAsPrettyJSON(label string, data interface{}) {
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Printf("%s: error marshaling JSON: %v", label, err)
		return
	}
	log.Printf("%s:\n%s", label, string(jsonBytes))
}

func VerifySignature(secret, signatureHeader string, body []byte) bool {
	// Remove any possible prefix like "sha256=" from signature header
	receivedSig := strings.TrimPrefix(signatureHeader, "sha256=")

	// Compute HMAC-SHA256
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(body)
	expectedSig := hex.EncodeToString(mac.Sum(nil))

	// Use constant-time comparison to avoid timing attacks
	return hmac.Equal([]byte(receivedSig), []byte(expectedSig))
}

func HandleWebhookPayload() gin.HandlerFunc {
	return func(c *gin.Context) {
		logAsPrettyJSON("Headers", c.Request.Header)
		bodyBytes, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		var payload map[string]interface{}
		if err := c.ShouldBindJSON(&payload); err != nil {
			log.Printf("JSON bind error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  "Invalid JSON payload",
				"detail": err.Error(),
			})
			return
		}
		logAsPrettyJSON("Parsed Payload", payload)
		c.JSON(http.StatusCreated, gin.H{
			"message": "Payload received successfully",
			"payload": payload,
		})
	}
}

func HandleGitHubPayload() gin.HandlerFunc {
	return func(c *gin.Context) {
		var secret = os.Getenv("GITHUB_WEBHOOK_SECRET")

		logAsPrettyJSON("Headers", c.Request.Header)

		bodyBytes, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		signature := c.GetHeader("X-Hub-Signature-256")
		if !VerifySignature(secret, signature, bodyBytes) {
			log.Println("Invalid signature")
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid signature",
			})
			return
		}

		var payload models.GitHubPushPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			log.Printf("JSON bind error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  "Invalid JSON payload",
				"detail": err.Error(),
			})
			return
		}

		logAsPrettyJSON("Parsed Payload", payload)

		c.JSON(http.StatusCreated, gin.H{
			"message":  "Payload received",
			"branch":   payload.Ref,
			"commits":  len(payload.Commits),
			"repo":     payload.Repository.FullName,
			"pusher":   payload.Pusher.Name,
			"head_sha": payload.HeadCommit.ID,
		})
	}
}
