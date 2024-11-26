package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hilton-james/FetchExercise/config"
	"github.com/hilton-james/FetchExercise/internal/core/entities"
	"github.com/hilton-james/FetchExercise/internal/ports"
	"github.com/hilton-james/FetchExercise/internal/usecases"
	"go.uber.org/zap"
)

const (
	ErrFailedToParsedJson = "failed to parse json"
	ErrInvalidInput       = "Invalid input"
	ErrNotFound           = "not found"
	ErrFailedCalculate    = "failed to calculate"
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

	receiptID, err := usecases.ProcessReceipt(r.repo, &input)
	if err != nil {
		HandleError(c, r.logger, http.StatusInternalServerError, "Failed to process receipt", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": receiptID}) // TODO: use struct
}

func (r *Receipt) GetPoints(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		HandleError(c, r.logger, http.StatusBadRequest, ErrInvalidInput, err)
		return
	}

	// TODO: Might need to have a cash mechanism before getting from the db
	receipt, err := r.repo.GetByID(id)
	if err != nil {
		HandleError(c, r.logger, http.StatusNotFound, ErrNotFound, err)
		return
	}

	points, err := usecases.CalculatePoints(receipt)
	if err != nil {
		HandleError(c, r.logger, http.StatusInternalServerError, ErrFailedCalculate, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"points": points}) // TODO: use struct
}
