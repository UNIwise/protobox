package proto

import (
	"errors"
	"io"
	"os"
	"os/exec"
	"path"

	"github.com/UNIwise/protobuf-cli/internal/docker"
)

func Generate(proto string, language string, out string, docker bool) error {
	os.MkdirAll(out, os.ModePerm)

	protoDest := path.Join(out, path.Base(proto))

	os.Rename(proto, protoDest)

	args, err := generateLanguageArgs(protoDest, language, out)

	if err != nil {
		return err
	}

	if docker {
		err = dockerGenerate(args)
	} else {
		err = localGenerate(args)
	}

	return err
}

func generateLanguageArgs(proto string, language string, out string) ([]string, error) {
	args := make([]string, 2)

	switch language {
	case "go":
		args[0] = "--go_out=plugins=grpc,import_path=" + out + ":."
	case "ts":
		args[0] = "--ts_out=service=true:."
	case "js":
		args[0] = "--js_out=" + out
	case "php":
		args[0] = "--php_out=" + out
	case "python":
		args[0] = "--python_out=."
	case "java":
		args[0] = "--java_out=."
	case "cpp":
		args[0] = "--cpp_out=."
	case "ruby":
		args[0] = "--ruby_out=."
	default:
		return nil, errors.New("Language \"" + language + "\" not supported")
	}

	args[1] = proto

	return args, nil
}

func localGenerate(args []string) error {
	cmd := exec.Command("protoc", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func dockerGenerate(args []string) error {
	cmd := "protoc"
	return docker.Run(cmd, "./", "test")
}

func copyFile(src string, dst string) error {
	_, err := os.Stat(src)
	if err != nil {
		return err
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}

	defer destination.Close()
	_, err = io.Copy(destination, source)

	return err
}
