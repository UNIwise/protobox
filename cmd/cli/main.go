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
	tempDir            = "/tmp/protobox"
	defaultDockerImage = "wiseflow/protobox"
	yamlFile           = "protobox.yaml"
)

var (
	dockerVar bool
)

func main() {
	generateCommand := flag.NewFlagSet("generate", flag.ExitOnError)
	generateCommand.BoolVar(&dockerVar, "docker", false, "Specifies if docker builder should be used.")

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "lint":
			lint()
			break
		case "generate":
			generateCommand.Parse(os.Args[2:])
			generate(dockerVar)
			break
		case "--help":
		case "help":
		default:
			flag.PrintDefaults()
		}
	} else {
		generateCommand.Parse(os.Args[2:])
		generate(dockerVar)
	}
}

func lint() {
	err := yaml.Lint(yamlFile)
	checkError(err)

	fmt.Println(color.GreenString("::"), "No errors, all ok")
}

func generate(dockerVar bool) {
	startTime := time.Now()

	def, err := yaml.ReadStruct(yamlFile)

	checkError(err)

	fmt.Println(color.BlueString("::"), len(def.Services), "protobuf service(s) found")

	if dockerVar {
		if def.Builder == "" {
			fmt.Println(color.BlueString("::"), "Using default docker builder", defaultDockerImage)
			def.Builder = defaultDockerImage
		} else {
			fmt.Println(color.BlueString("::"), "Using docker builder", def.Builder)
		}
	}

	cleanTempDir()

	fmt.Println()

	for _, s := range def.Services {

		if s.Repo != "" {
			fmt.Println(color.CyanString("==>"), rightPad("Syncing:", " ", 20), s.Repo, s.Branch, s.Commit)

			_, err := git.Clone(s.Repo, s.Branch, tempDir)
			checkError(err)
		}

		for _, l := range s.Out {
			fmt.Println(color.CyanString("==>"), rightPad("Generating ["+l.Language+"]:", " ", 20), path.Base(s.Proto))

			protoFile := s.Proto
			if s.Repo != "" {
				protoFile = path.Join(tempDir, s.Proto)
			}

			err := proto.Generate(protoFile, l.Language, l.Path, dockerVar, def.Builder)
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
