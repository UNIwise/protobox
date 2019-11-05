package proto

import (
	"errors"
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/UNIwise/protobox/internal/docker"
)

func Generate(proto string, language string, src string, out string, local bool, dockerImage string) error {
	os.MkdirAll(out, os.ModePerm)

	protoOut := path.Join(out, path.Base(proto))

	if proto != protoOut {
		err := copyFile(proto, protoOut)
		if err != nil {
			return err
		}
	}

	if language == "none" {
		return nil
	}

	outAbs, err := filepath.Abs(out)
	if err != nil {
		return err
	}

	srcAbs, err := filepath.Abs(src)
	if err != nil {
		return err
	}

	if local {
		return localGenerate(proto, language, srcAbs, outAbs)
	} else {
		return dockerGenerate(dockerImage, proto, language, srcAbs, outAbs)
	}
}

func generateLanguageArgs(proto string, language string, out string) ([]string, error) {
	args := make([]string, 3)

	args[0] = "--proto_path=" + path.Dir(proto)

	switch language {
	case "go":
		args[1] = "--go_out=plugins=grpc:" + out
	case "ts":
		args[1] = "--ts_out=" + out
	case "js":
		args[1] = "--js_out=" + out
	case "php":
		args[1] = "--php_out=" + out
	case "python":
		args[1] = "--python_out=" + out
	case "java":
		args[1] = "--java_out=" + out
	case "cpp":
		args[1] = "--cpp_out=" + out
	case "ruby":
		args[1] = "--ruby_out=" + out
	default:
		return nil, errors.New("Language \"" + language + "\" not supported")
	}

	args[2] = proto

	return args, nil
}

func localGenerate(proto string, language string, src string, out string) error {
	if !HasProtoc() {
		return errors.New("No protoc binary found")
	}

	args, err := generateLanguageArgs(path.Join(src, proto), language, out)
	if err != nil {
		return err
	}

	cmd := exec.Command("protoc", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func dockerGenerate(dockerImage string, proto string, language string, src string, out string) error {
	args, err := generateLanguageArgs(proto, language, "/out")
	if err != nil {
		return err
	}

	return docker.Run(dockerImage, "protoc", args, src+":/mnt", out+":/out")
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
