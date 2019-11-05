package yaml

import "github.com/UNIwise/protobox/internal/defaults"

// Definition describe the configuration yaml file
type Definition struct {
	Syntax   string    `yaml:"syntax"validate:"required,syntax"`
	Builder  string    `yaml:"builder,omitempty"`
	Services []Service `yaml:"services"validate:"dive"`
}

type Service struct {
	Repo   string `yaml:"repo,omitempty"`
	Proto  string `yaml:"proto"validate:"required"`
	Commit string `yaml:"commit,omitempty"`
	Branch string `yaml:"branch,omitempty"`
	Out    []Out  `yaml:"out"validate:"required,dive"`
}

type Out struct {
	Path     string `yaml:"path"validate:"required"`
	Language string `yaml:"language"validate:"required,language"`
}

func New() *Definition {
	d := Definition{}
	d.Syntax = defaults.Syntax
	d.Builder = defaults.DockerImage

	return &d
}
