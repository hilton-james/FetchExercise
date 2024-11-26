package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hilton-james/FetchExercise/config"
	"go.uber.org/zap"
)

type Receipt struct {
	logger *zap.Logger
	cfg    config.Receipt
}

func NewReceipt(cfg config.Receipt, logger *zap.Logger) *Receipt {
	return &Receipt{
		logger: logger,
		cfg:    cfg,
	}
}

func (r *Receipt) Register(router *gin.RouterGroup) {
	router.POST("/process", r.ProcessReceipt)
	router.GET("/:id/points", r.GetPoints)
}

func (r *Receipt) ProcessReceipt(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ProcessReceipt"})
}

func (r *Receipt) GetPoints(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetPoints"})
}
