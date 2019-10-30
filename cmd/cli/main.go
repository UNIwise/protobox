package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"

	git "github.com/UNIwise/protobuf-cli/internal/git"
	proto "github.com/UNIwise/protobuf-cli/internal/proto"
	yaml "github.com/UNIwise/protobuf-cli/internal/yaml"

	color "github.com/fatih/color"
)

const (
	tempDir = "./.proto"
)

var (
	configVar string
)

func main() {
	lintCommand := flag.NewFlagSet("lint", flag.ExitOnError)
	lintCommand.StringVar(&configVar, "config", "protobuf.yaml", "Define a config file to use.")

	generateCommand := flag.NewFlagSet("install", flag.ExitOnError)
	generateCommand.StringVar(&configVar, "config", "protobuf.yaml", "Define a config file to use.")

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "lint":
			lintCommand.Parse(os.Args[2:])
			yaml.Lint(configVar)
			break
		case "generate":
			generateCommand.Parse(os.Args[2:])
			generate(configVar)
		case "--help":
		case "help":
		default:
			flag.PrintDefaults()
		}
	} else {
		generate(configVar)
	}
}

func generate(config string) {
	startTime := time.Now()

	yaml.Lint(config)

	def := yaml.ReadStruct(config)

	fmt.Println(color.BlueString("::"), len(def.Services), "protobuf service(s) found")

	cleanTempDir()

	for _, s := range def.Services {
		fmt.Println(color.CyanString("\n==>"), "Syncing:\t", s.Repo, s.Branch, s.Commit)

		_, err := git.Clone(s.Repo, s.Branch, tempDir)

		checkError(err)

		for _, l := range s.Out {
			fmt.Println(color.CyanString("==>"), "Generating:\t", s.Proto)
			err := proto.Generate(path.Join(tempDir, s.Proto), l.Language, l.Path)

			checkError(err)
		}

		cleanTempDir()
	}

	os.Remove(tempDir)

	duration := time.Now().Sub(startTime) / time.Nanosecond

	fmt.Println(color.GreenString("\n::"), "Done", duration)
}

func cleanTempDir() {
	dir, err := ioutil.ReadDir(tempDir)

	if err != nil {
		return
	}

	for _, d := range dir {
		os.RemoveAll(path.Join([]string{tempDir, d.Name()}...))
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(color.RedString("==>"), err)
		os.Exit(1)
	}
}
