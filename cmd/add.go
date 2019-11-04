package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add [URL] [LANGUAGE] [OUT DIR]",
	Short: "Add proto dependency to protobox.yaml",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		add(args[0], args[1], args[2])
	},
}

func add(url string, language string, out string) {
	if strings.Contains(url, "/blob/") {
		checkError(addGitProto(url, language))
	}

	checkError(errors.New("Unknown resource"))
}

func addGitProto(url string, language string) error {
	res := strings.Split(url, "/blob/")
	idx := strings.Index(res[1], "/")

	if len(res) < 2 || idx == -1 {
		return errors.New("Unknown resource")
	}

	repo := res[0]
	ref := res[1][:idx]
	file := res[1][idx+1:]

	fmt.Println(repo, ref, file)

	return nil
}
