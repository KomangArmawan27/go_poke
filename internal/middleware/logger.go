// middleware/logger.go
package middleware

import (
	"bytes"
	"io"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go-api/config"
	"go-api/internal/models"
)

// Custom response writer to capture response body
type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b) // Capture the response
	return w.ResponseWriter.Write(b) // Write response as normal
}

// LoggerMiddleware logs and stores request/response details
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Read request body
		reqBody, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))

		// Wrap response writer
		bw := &bodyWriter{body: new(bytes.Buffer), ResponseWriter: c.Writer}
		c.Writer = bw

		c.Next()
		duration := time.Since(start)
		

		// Collect data
		method := c.Request.Method
		path := c.Request.URL.Path
		status := c.Writer.Status()
		ip := c.ClientIP()

		// Log to console
		log.Printf(
			"%s | %s | %s | %d | %v",
			method, path, ip, status, duration,
		)

		logEntry := models.Log{
			Method:       c.Request.Method,
			URI:          c.Request.RequestURI,
			ClientIP:     c.ClientIP(),
			StatusCode:   c.Writer.Status(),
			Duration:     duration.String(),
			RequestBody:  string(reqBody),
			ResponseBody: bw.body.String(),
			CreatedAt:    time.Now(),
		}

		config.DB.Create(&logEntry)
	}
}
