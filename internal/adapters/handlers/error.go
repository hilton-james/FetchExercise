package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func HandleError(c *gin.Context, logger *zap.Logger, statusCode int, errorMessage string, err error) {
	logger.Error(errorMessage, zap.Error(err))
	c.JSON(statusCode, gin.H{"error": errorMessage}) // TODO: use struct for error
}
