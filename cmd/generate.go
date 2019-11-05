package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/UNIwise/protobox/internal/defaults"
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

	def, err := yaml.Read(defaults.YamlFile)

	checkError(err)

	fmt.Println(color.BlueString("::"), len(def.Services), "protobuf service(s) found")

	if !localBin {
		if def.Builder == "" {
			fmt.Println(color.BlueString("::"), "Using default docker builder", defaults.DockerImage)
			def.Builder = defaults.DockerImage
		} else {
			fmt.Println(color.BlueString("::"), "Using docker builder", def.Builder)
		}
	}

	cleanTempDir()

	fmt.Println()

	for _, s := range def.Services {
		srcDir := "./"

		if s.Repo != "" {
			fmt.Println(color.CyanString("==>"), rightPad("Syncing:", " ", 20), s.Repo, s.Ref)

			err := git.Clone(s.Repo, s.Ref, defaults.TempDir)
			checkError(err)

			srcDir = defaults.TempDir
		}

		for _, l := range s.Out {
			fmt.Println(color.CyanString("==>"), rightPad("Generating ["+l.Language+"]:", " ", 20), path.Base(s.Proto))

			err := proto.Generate(s.Proto, l.Language, srcDir, l.Path, localBin, def.Builder, s.Repo != "")
			checkError(err)
		}

		cleanTempDir()
	}

	os.Remove(defaults.TempDir)

	duration := time.Now().Sub(startTime) / time.Nanosecond
	fmt.Println(color.GreenString("\n::"), "Done", duration)
}

func cleanTempDir() {
	dir, err := ioutil.ReadDir(defaults.TempDir)
	if err != nil {
		return
	}

	for _, d := range dir {
		os.RemoveAll(path.Join([]string{defaults.TempDir, d.Name()}...))
	}
}
