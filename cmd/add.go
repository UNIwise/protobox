package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add proto dependency to protobox.yaml",
	Run: func(cmd *cobra.Command, args []string) {
		add()
	},
}

func add() {
	checkError(errors.New("Command not implemented"))
}
