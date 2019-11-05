package cmd

import (
	"errors"
	"os"
	"strings"

	"github.com/UNIwise/protobox/internal/defaults"
	"github.com/UNIwise/protobox/internal/yaml"
	"github.com/spf13/cobra"
)

var errUnknownResource = errors.New("Unknown resource")

var addCmd = &cobra.Command{
	Use:   "add [URI] [LANGUAGE] [OUT DIR]",
	Short: "Add proto dependency to " + defaults.YamlFile,
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		add(args[0], args[1], args[2])
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func add(uri string, language string, out string) {
	var d *yaml.Definition
	var err error

	if !yaml.Exists() {
		d = yaml.New()
	} else {
		d, err = yaml.Load()
		checkError(err)
	}

	if strings.Contains(uri, "http") && strings.Contains(uri, "/blob/") {
		checkError(addGitProto(d, uri, language, out))
	} else if fileExists(uri) {
		addLocalProto(d, uri, language, out)
	} else {
		checkError(errUnknownResource)
	}
}

func addGitProto(d *yaml.Definition, uri string, language string, out string) error {
	res := strings.Split(uri, "/blob/")
	idx := strings.Index(res[1], "/")

	if len(res) < 2 || idx == -1 {
		return errUnknownResource
	}

	repo := res[0]
	ref := res[1][:idx]
	file := res[1][idx+1:]

	outStruct := []yaml.Out{
		yaml.Out{
			Language: language,
			Path:     out,
		},
	}

	return addService(d, yaml.Service{
		Repo:   repo,
		Branch: ref,
		Proto:  file,
		Out:    outStruct,
	})
}

func addLocalProto(d *yaml.Definition, uri string, language string, out string) error {
	outStruct := []yaml.Out{
		yaml.Out{
			Language: language,
			Path:     out,
		},
	}

	return addService(d, yaml.Service{
		Proto: uri,
		Out:   outStruct,
	})
}

func addService(d *yaml.Definition, s yaml.Service) error {
	d.Services = append(d.Services, s)
	return yaml.Save(d)
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
