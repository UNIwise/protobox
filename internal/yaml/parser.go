package yaml

import (
	"errors"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func Read(path string) (*Definition, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.New("no protobox.yaml file found")
	}

	t := Definition{}

	err = yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		return nil, err
	}

	err = Validate(t)

	return &t, err
}

func Write(path string, def Definition) error {
	bytes, err := yaml.Marshal(def)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, bytes, os.ModePerm)
}

func Lint(path string) error {
	_, err := Read(path)
	return err
}
