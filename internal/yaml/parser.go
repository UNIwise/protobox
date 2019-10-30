package yaml

import (
	"errors"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

func ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func ReadStruct(path string) (*Definition, error) {
	data, err := ReadFile(path)
	if err != nil {
		return nil, errors.New("no protobuf.yaml file found")
	}

	t := Definition{}

	err = yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		return nil, err
	}

	err = Validate(t)

	return &t, err
}

func Lint(path string) error {
	_, err := ReadStruct(path)
	return err
}
