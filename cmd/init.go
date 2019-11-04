package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes protobox.yaml",
	Run: func(cmd *cobra.Command, args []string) {
		initialize()
	},
}

func initialize() {
	checkError(errors.New("Command not implemented"))
}
