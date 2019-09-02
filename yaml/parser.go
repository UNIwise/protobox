package yaml

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func ReadStruct(path string) *Definition {
	data, err := ReadFile(path)

	if err != nil {
		log.Fatalln("No protobuf.yaml file found.")
		return nil
	}

	t := Definition{}

	err = yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return &t
}

func Lint(path string) {
	ReadStruct(path)
}
