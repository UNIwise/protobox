package cmd

import (
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

func init() {

}
