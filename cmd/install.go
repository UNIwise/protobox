package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/UNIwise/protobox/internal/git"
	"github.com/UNIwise/protobox/internal/proto"
	"github.com/UNIwise/protobox/internal/yaml"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	localBin bool
)

func init() {
	genreateCmd.PersistentFlags().BoolVar(&localBin, "local", false, "Specifies if local protoc should be used")

	rootCmd.AddCommand(genreateCmd)
}

var genreateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate protobuf source files",
	Run: func(cmd *cobra.Command, args []string) {
		generate()
	},
}

func generate() {
	startTime := time.Now()

	def, err := yaml.ReadStruct(yamlFile)

	checkError(err)

	fmt.Println(color.BlueString("::"), len(def.Services), "protobuf service(s) found")

	if !localBin {
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

			if s.Branch != "" {
				s.Branch = "refs/heads/" + s.Branch
			}

			_, err := git.Clone(s.Repo, s.Branch, s.Commit, tempDir)
			checkError(err)
		}

		for _, l := range s.Out {
			fmt.Println(color.CyanString("==>"), rightPad("Generating ["+l.Language+"]:", " ", 20), path.Base(s.Proto))

			protoFile := s.Proto
			if s.Repo != "" {
				protoFile = path.Join(tempDir, s.Proto)
			}

			err := proto.Generate(protoFile, l.Language, l.Path, localBin, def.Builder)
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
