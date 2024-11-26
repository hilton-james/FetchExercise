package server

import (
	"log"

	"github.com/hilton-james/FetchExercise/pkg/logger"
	"github.com/spf13/cobra"
)

func main(_ *cobra.Command, _ []string) {
	logger, cancel, err := logger.New(true) // TODO: Should be gotten from .env
	if err != nil {
		log.Fatalf("logger failed %s", err)
	}

	defer cancel()

	logger.Info("server is running")

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
