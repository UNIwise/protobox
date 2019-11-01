package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	docker "github.com/UNIwise/protobox/internal/docker"
	proto "github.com/UNIwise/protobox/internal/proto"
)

type ServiceCheck struct {
	Name     string
	Func     func() bool
	Optional bool
}

func init() {
	rootCmd.AddCommand(checkCmd)
}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check if your system have the needed dependencies for protobox to work",
	Run: func(cmd *cobra.Command, args []string) {
		check()
	},
}

var checks = []ServiceCheck{
	{
		Name:     "Docker",
		Func:     docker.HasDocker,
		Optional: false,
	},
	{
		Name:     "protoc",
		Func:     proto.HasProtoc,
		Optional: true,
	},
	{
		Name:     "protoc-gen-go",
		Func:     proto.HasProtocGo,
		Optional: true,
	},
	{
		Name:     "protoc-gen-ts",
		Func:     proto.HasProtocTypescript,
		Optional: true,
	},
}

func check() {
	fmt.Println(color.BlueString("::"), "Checking system dependencies...\n")

	failsTotal := 0
	failsFatal := 0

	for _, c := range checks {
		ok := c.Func()

		res := color.GreenString("OK")
		if !ok {
			failsTotal++

			res = color.RedString("MISSING")

			if c.Optional {
				res += " (optional)"
			} else {
				res += " (required)"
				failsFatal++
			}
		}

		fmt.Println(color.CyanString("==>"), rightPad(c.Name, " ", 20), res)
	}

	fmt.Println()
	if failsFatal > 0 {
		fmt.Println(color.RedString("::"), "A required dependency is missing, without it protobox will not work")
	} else if failsTotal > 0 {
		fmt.Println(color.YellowString("::"), "An optional dependency is missing, use the docker builder to insure a complete build enviroment")
	} else {
		fmt.Println(color.GreenString("::"), "All dependencies are installed")
	}
}
