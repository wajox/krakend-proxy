package cli

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var RootCmd = cobra.Command{}

func RegisterCommands() {
	ServeCmd.Flags().Bool("jsonlogs", false, "Use JSON for logs")
	RootCmd.AddCommand(ServeCmd)
}

func Execute() {
	RegisterCommands()

	if err := RootCmd.Execute(); err != nil {
		log.Fatal().Err(err)
	}
}
