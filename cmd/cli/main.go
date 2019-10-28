package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	git "github.com/UNIwise/protobuf-cli/internal/git"
	proto "github.com/UNIwise/protobuf-cli/internal/proto"
	yaml "github.com/UNIwise/protobuf-cli/internal/yaml"

	color "github.com/fatih/color"
)

var (
	configVar string
)

func main() {

	lintCommand := flag.NewFlagSet("lint", flag.ExitOnError)
	lintCommand.StringVar(&configVar, "config", "protobuf.yaml", "Define a config file to use.")

	installCommand := flag.NewFlagSet("install", flag.ExitOnError)
	installCommand.StringVar(&configVar, "config", "protobuf.yaml", "Define a config file to use.")

	if len(os.Args) > 1 {

		switch os.Args[1] {
		case "lint":
			lintCommand.Parse(os.Args[2:])
			yaml.Lint(configVar)
			break
		case "install":
			installCommand.Parse(os.Args[2:])
			install(configVar)
		case "--help":
		case "help":
		default:
			flag.PrintDefaults()
		}
	} else {
		// TODO
	}
}

func install(config string) {
	startTime := time.Now()

	yaml.Lint(config)

	def := yaml.ReadStruct(config)

	fmt.Println(color.BlueString("::"), len(def.Services), "protobuf client to generate")

	for _, s := range def.Services {
		fmt.Println(color.GreenString("==>"), "syncing", s.Repo, s.Branch, s.Commit)

		_, err := git.Clone(s.Repo, s.Branch, "./tmp")

		if err != nil {
			fmt.Println(color.RedString("==>"), err)
			os.Exit(1)
		}

		for _, l := range s.Out {
			fmt.Println(color.GreenString("==>"), "generating", s.Proto, "->", l.Path)
			proto.Generate(s.Proto, l.Language, l.Path)
		}
	}

	duration := time.Now().Sub(startTime) / time.Nanosecond

	fmt.Println(color.GreenString("::"), "Done", duration)

}
