package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hilton-james/FetchExercise/config"
	"github.com/hilton-james/FetchExercise/internal/adapters/handlers"
	"github.com/hilton-james/FetchExercise/pkg/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func main(_ *cobra.Command, _ []string) {
	cfg, err := config.NewReceipt()
	if err != nil {
		log.Fatalf("failed to load config %s", err)
	}

	logger, cancel, err := logger.New(cfg) // TODO: Should be gotten from .env
	if err != nil {
		log.Fatalf("logger failed %s", err)
	}
	defer cancel()

	handler := gin.Default()
	receiptHandler := handlers.NewReceipt(cfg, logger.Named("receiptHandler"))

	receiptHandler.Register(handler.Group("/receipts"))

	logger.Info("server is running")

	// TODO: health check endpoint should be implemented (/heath)

	// TODO: server should be shutdown properly (SIGINT or SIGTERM).
	if err := handler.Run(cfg.Port); err != nil {
		logger.Fatal("failed to run server", zap.Error(err))
	}
}

// Register server command.
func Register(
	root *cobra.Command,
) {
	root.AddCommand(
		//nolint: exhaustruct
		&cobra.Command{
			Use:   "server",
			Short: "Run server to serve the requests",
			Run:   main,
		},
	)
}
