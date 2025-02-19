package zaplogger

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ZapLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// End timer
		latency := time.Since(start)

		// Get status code
		status := c.Writer.Status()

		// Log request details
		logger.Debug("Request",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", status),
			zap.String("client_ip", c.ClientIP()),
			zap.Duration("latency", latency),
			zap.String("user_agent", c.Request.UserAgent()),
		)

		// Log any errors (if present)
		if len(c.Errors) > 0 {
			zapErrors := stringsToErrors(c.Errors.Errors())
			logger.Error("Request Errors",
				zap.Errors("errors", zapErrors),
			)
		}
	}
}

func stringsToErrors(strings []string) []error {
	errorsList := make([]error, len(strings))
	for i, s := range strings {
		errorsList[i] = errors.New(s)
	}
	return errorsList
}
