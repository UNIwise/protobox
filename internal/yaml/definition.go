package yaml

// Definition describe the configuration yaml file 'protobox.yaml'
type Definition struct {
	Services []Service `json:"services"validate:"dive"`
	Builder  string    `json:"builder,omitempty"`
	Syntax   string    `json:"syntax"validate:"required,syntax"`
}

type Service struct {
	Repo   string `json:"repo"`
	Proto  string `json:"proto"validate:"required"`
	Commit string `json:"commit,omitempty"`
	Branch string `json:"branch,omitempty"`
	Out    []Out  `json:"out"validate:"required,dive"`
}

type Out struct {
	Path     string `json:"path"validate:"required"`
	Language string `json:"language"validate:"required,language"`
}
