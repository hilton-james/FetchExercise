package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hilton-james/FetchExercise/config"
	"github.com/hilton-james/FetchExercise/internal/core/entities"
	"github.com/hilton-james/FetchExercise/internal/ports"
	"go.uber.org/zap"
)

const (
	ErrFailedToParsedJson = "failed to parse json"
	ErrInvalidInput       = "Invalid input"
)

type Receipt struct {
	logger *zap.Logger
	cfg    config.Receipt
	repo   ports.Repository
}

func NewReceipt(cfg config.Receipt, logger *zap.Logger, repo ports.Repository) *Receipt {
	return &Receipt{
		logger: logger,
		cfg:    cfg,
		repo:   repo,
	}
}

func (r *Receipt) Register(router *gin.RouterGroup) {
	router.POST("/process", r.ProcessReceipt)
	router.GET("/:id/points", r.GetPoints)
}

func (r *Receipt) ProcessReceipt(c *gin.Context) {
	var input entities.Receipt
	if err := c.ShouldBindJSON(&input); err != nil {
		HandleError(c, r.logger, http.StatusBadRequest, ErrFailedToParsedJson, err)
		return
	}

	if err := input.Validate(); err != nil {
		HandleError(c, r.logger, http.StatusBadRequest, ErrInvalidInput, err)
		return
	}

	input.ID = uuid.New().String() // TODO: consider expiration time if need be

	// TODO: here we can check if it's duplicated
	if err := r.repo.Save(&input); err != nil {
		HandleError(c, r.logger, http.StatusInternalServerError, "Failed to save receipt", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": input.ID})
}

func (r *Receipt) GetPoints(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetPoints"})
}
