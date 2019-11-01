package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const (
	tempDir            = "/tmp/protobox"
	defaultDockerImage = "wiseflow/protobox"
	yamlFile           = "protobox.yaml"
	version            = "1.0.0"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "protobox",
		Short: "gRPC source generation and dependency management like a good boy.",
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func checkError(err error) {
	if err != nil {
		fmt.Println(color.RedString("==>"), err)
		os.Exit(1)
	}
}

func rightPad(input string, padChar string, length int) string {
	result := input

	add := length - len(input)

	for i := 0; i < add; i++ {
		result = result + padChar
	}

	return result
}
