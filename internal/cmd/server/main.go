package server

import (
	"log"

	"github.com/spf13/cobra"
)

func main(_ *cobra.Command, _ []string) {
	log.Println("welcome to our server")
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
