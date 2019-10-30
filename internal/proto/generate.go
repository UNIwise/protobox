package proto

import (
	"errors"
	"os"
	"os/exec"
	"path"
	"strings"
)

func Generate(proto string, language string, out string) error {
	var err error

	if hasLocalBinaries() {
		err = localGenerate(proto, language, out)
	} else {
		err = dockerGenerate(proto, language, out)
	}

	return err
}

func hasLocalBinaries() bool {
	return true
}

func generateLanguageArgs(proto string, language string, out string) ([]string, error) {
	args := make([]string, 2)

	switch language {
	case "go":
		args[0] = "--go_out=plugins=grpc,import_path=" + out + ":."
	default:
		return nil, errors.New("Language \"" + language + "\" not supported")
	}

	args[1] = proto

	return args, nil
}

func localGenerate(proto string, language string, out string) error {
	os.MkdirAll(out, os.ModePerm)

	protoDest := path.Join(out, path.Base(proto))

	os.Rename(proto, protoDest)

	args, err := generateLanguageArgs(protoDest, language, out)

	if err != nil {
		return err
	}

	cmd := exec.Command("protoc", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func copyGeneratedFile(proto string, out string) {
	genDir := path.Dir(proto)
	filenameExt := path.Base(proto)
	filename := strings.TrimSuffix(filenameExt, path.Ext(filenameExt))

	os.Rename(path.Join(genDir, filename)+".pb.go", path.Join(out, filename)+".pb.go")
}

func dockerGenerate(proto string, language string, out string) error {
	return nil
}
