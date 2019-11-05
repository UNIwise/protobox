package cmd

import (
	"errors"
	"fmt"

	"github.com/UNIwise/protobox/internal/defaults"
	"github.com/UNIwise/protobox/internal/yaml"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize " + defaults.YamlFile,
	Run: func(cmd *cobra.Command, args []string) {
		initialize()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func initialize() {
	if yaml.Exists() {
		checkError(errors.New(defaults.YamlFile + " already exists"))
		return
	}

	err := yaml.Save(yaml.New())
	checkError(err)

	fmt.Println(color.GreenString("::"), defaults.YamlFile, "initialized 🎉")
}
