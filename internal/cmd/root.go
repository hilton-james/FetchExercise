package cmd

import (
	"log"

	"github.com/hilton-james/FetchExercise/internal/cmd/server"
	"github.com/spf13/cobra"
)

func Execute() {

	//nolint: exhaustruct
	root := &cobra.Command{
		Use:     "server",
		Short:   "Receipt Server which serve Take-Home Exercise",
		Version: "v0.0.1",
	}

	server.Register(root)

	if err := root.Execute(); err != nil {
		log.Fatalf("failed to execute root command %s", err)
	}
}
