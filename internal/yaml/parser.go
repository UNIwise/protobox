package yaml

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/UNIwise/protobox/internal/defaults"
	yaml "gopkg.in/yaml.v2"
)

func Read(path string) (*Definition, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.New("no " + defaults.YamlFile + " file found")
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

	return ioutil.WriteFile(path, bytes, 0664)
}

func Lint(path string) error {
	_, err := Read(path)
	return err
}

func Exists() bool {
	if _, err := os.Stat(defaults.YamlFile); os.IsNotExist(err) {
		return false
	}
	return true
}

func Load() (*Definition, error) {
	return Read(defaults.YamlFile)
}

func Save(def *Definition) error {
	return Write(defaults.YamlFile, *def)
}
