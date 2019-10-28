package proto

import (
	"os/exec"
)

func Generate(proto string, language string, out string) {
	if hasLocalBinaries() {
		localGenerate(proto, language, out)
	} else {
		dockerGenerate(proto, language, out)
	}
}

func hasLocalBinaries() bool {
	return true
}

func generateLanguageArgs(proto string, language string, out string) []string {
	args := make([]string, 2)

	args[0] = "--go_out=plugins=grpc,import_path=out/:out/"
	args[1] = "tmp/" + proto

	return args
}

func localGenerate(proto string, language string, out string) error {
	args := generateLanguageArgs(proto, language, out)
	cmd := exec.Command("protoc", args...)
	return cmd.Run()
}

func dockerGenerate(proto string, language string, out string) {

}
