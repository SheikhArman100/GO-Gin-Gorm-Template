package response

import (
	"runtime/debug"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"my-project/internal/logger"
)

type ErrorResponse struct {
	Success       bool        `json:"success"`
	StatusCode    int         `json:"statusCode"`
	Message       string      `json:"message"`
	ErrorMessages interface{} `json:"errorMessages,omitempty"`
	Stack         string      `json:"stack,omitempty"`
}

// ApiError sends an immediate error response and aborts further processing
func ApiError(c *gin.Context, statusCode int, message string, errorMessages ...interface{}) {
	
	stack := ""
	if os.Getenv("ENV") != "production" {
		stack = string(debug.Stack())
	}

	errResp := ErrorResponse{
		Success:    false,
		StatusCode: statusCode,
		Message:    message,
	}

	if len(errorMessages) > 0 {
		errResp.ErrorMessages = errorMessages[0]
	}

	if stack != "" {
		errResp.Stack = stack
	}

	// Log the error using your custom errorLogger
	logger.ErrorLogger.Error("API error occurred",
		zap.Int("statusCode", statusCode),
		zap.String("message", message),
		zap.Any("errorMessages", errResp.ErrorMessages),
		// zap.String("stack", stack),
		zap.String("path", c.Request.URL.Path),
		zap.String("method", c.Request.Method),
		zap.String("clientIP", c.ClientIP()),
	)

	c.AbortWithStatusJSON(statusCode, gin.H{"error": errResp})
}
