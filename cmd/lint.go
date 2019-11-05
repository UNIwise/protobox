package cmd

import (
	"fmt"

	"github.com/UNIwise/protobox/internal/defaults"
	"github.com/UNIwise/protobox/internal/yaml"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(lintCmd)
}

var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Lint syntax checks your " + defaults.YamlFile,
	Run: func(cmd *cobra.Command, args []string) {
		lint()
	},
}

func lint() {
	err := yaml.Lint(defaults.YamlFile)
	checkError(err)

	fmt.Println(color.GreenString("::"), "No errors, all ok 👍")
}
