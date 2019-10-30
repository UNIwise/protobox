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
	dockerVar bool
)

func main() {
	lintCommand := flag.NewFlagSet("lint", flag.ExitOnError)
	lintCommand.StringVar(&configVar, "config", "protobuf.yaml", "Specifies a config file to use.")

	generateCommand := flag.NewFlagSet("generate", flag.ExitOnError)
	generateCommand.StringVar(&configVar, "config", "protobuf.yaml", "Specifies a config file to use.")
	generateCommand.BoolVar(&dockerVar, "docker", false, "Specifies if docker builder should be used.")

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "lint":
			lintCommand.Parse(os.Args[2:])
			lint(configVar)
			break
		case "generate":
			generateCommand.Parse(os.Args[2:])
			generate(configVar, dockerVar)
			break
		case "--help":
		case "help":
		default:
			flag.PrintDefaults()
		}
	} else {
		generate(configVar, dockerVar)
	}
}

func lint(config string) {
	err := yaml.Lint(config)
	checkError(err)
}

func generate(config string, dockerVar bool) {
	startTime := time.Now()

	def, err := yaml.ReadStruct(config)

	checkError(err)

	fmt.Println(color.BlueString("::"), len(def.Services), "protobuf service(s) found")

	if dockerVar {
		fmt.Println(color.BlueString("::"), "Using docker builder", def.Builder)
	}

	cleanTempDir()

	for _, s := range def.Services {
		fmt.Println(color.CyanString("\n==>"), rightPad("Syncing:", " ", 20), s.Repo, s.Branch, s.Commit)

		_, err := git.Clone(s.Repo, s.Branch, tempDir)

		checkError(err)

		for _, l := range s.Out {
			fmt.Println(color.CyanString("==>"), rightPad("Generating ["+l.Language+"]:", " ", 20), s.Proto)
			err := proto.Generate(path.Join(tempDir, s.Proto), l.Language, l.Path, dockerVar)

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

func rightPad(input string, padChar string, length int) string {
	result := input

	add := length - len(input)

	for i := 0; i < add; i++ {
		result = result + padChar
	}

	return result
}
