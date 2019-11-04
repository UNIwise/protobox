package cmd

import (
	"fmt"

	"github.com/UNIwise/protobox/internal/defaults"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of protobox",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(defaults.Version)
	},
}
