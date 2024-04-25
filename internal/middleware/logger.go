package middleware

import (
	"bytes"
	"time"

	"gin-app/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// CustomResponseWriter is a wrapper around http.ResponseWriter
type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Wrap the response writer
		blw := &CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		// Process request
		c.Next()

		// Collect response data
		responseBody := blw.body.String()

		// Ensure Body can be read multiple times
		if err := c.Request.ParseForm(); err != nil {
			logger.Log.Errorf("ParseForm error: %v", err)
		}

		logger.Log.WithFields(logrus.Fields{
			"status_code": c.Writer.Status(),
			"latency":     time.Since(start),
			"client_ip":   c.ClientIP(),
			"method":      c.Request.Method,
			"path":        c.Request.URL.Path,
			"params":      c.Request.URL.Query(),
			"error":       c.Errors.ByType(gin.ErrorTypePrivate).String(),
			"user-agent":  c.Request.UserAgent(),
			"response":    responseBody,
		}).Info("Request details: ")
	}
}
