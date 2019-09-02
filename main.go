package main

import (
	"flag"
	"os"

	yaml "./yaml"
)

var flagvar string

func main() {

	lintCommand := flag.NewFlagSet("lint", flag.ExitOnError)
	lintCommand.StringVar(&flagvar, "config", "protobuf.yaml", "Define a config file to use.")

	if len(os.Args) > 1 {

		switch os.Args[1] {
		case "lint":
			lintCommand.Parse(os.Args[2:])
			yaml.Lint(flagvar)
			break
		case "--help":
		case "help":
		default:
			flag.PrintDefaults()
		}
	} else {
		// TODO
	}
}
