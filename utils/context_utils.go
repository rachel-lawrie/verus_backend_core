package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rachel-lawrie/verus_backend_core/zaplogger"
	"go.uber.org/zap"
)

func GetClientIDFromContext(c *gin.Context) (string, error) {
	logger := zaplogger.GetLogger()
	clientID, exists := c.Get("client_id")
	if !exists {
		logger.Error("Client ID not found in context")
		return "", fmt.Errorf("client_id not found")
	}
	logger.Debug("Client ID", zap.Any("client_id", clientID))
	clientIDStr, ok := clientID.(string)
	if !ok {
		logger.Error("Client ID is not a string")
		return "", fmt.Errorf("client_id is not a string")
	}
	return clientIDStr, nil
}
