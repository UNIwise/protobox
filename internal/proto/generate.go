package proto

import (
	"errors"
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/UNIwise/protobox/internal/defaults"
	"github.com/UNIwise/protobox/internal/docker"
)

func Generate(proto string, language string, src string, out string, local bool, dockerImage string, isCloned bool) error {
	os.MkdirAll(out, os.ModePerm)

	protoOut := path.Join(out, path.Base(proto))

	if isCloned || proto != protoOut {
		err := copyFile(path.Join(defaults.TempDir, proto), protoOut)
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
	args := []string{}

	args = append(args, "--proto_path="+path.Dir(proto))

	switch language {
	case "go":
		args = append(args, "--go_out=plugins=grpc:"+out)
	case "node":
		args = append(args, "--js_out=import_style=commonjs,binary:"+out)
		args = append(args, "--ts_out=service=grpc-node:"+out)
		args = append(args, "--grpc_out="+out)
	case "web":
		args = append(args, "--js_out=import_style=commonjs,binary:"+out)
		args = append(args, "--ts_out=service=grpc-web:"+out)
	case "php":
		args = append(args, "--php_out="+out)
		args = append(args, "--grpc_out="+out)
	case "python":
		args = append(args, "--python_out="+out)
	case "java":
		args = append(args, "--java_out="+out)
	case "cpp":
		args = append(args, "--cpp_out="+out)
	case "ruby":
		args = append(args, "--ruby_out="+out)
	default:
		return nil, errors.New("Language \"" + language + "\" not supported")
	}

	args = append(args, proto)

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
