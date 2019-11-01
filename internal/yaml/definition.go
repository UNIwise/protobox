package yaml

// Definition describe the configuration yaml file 'protobox.yaml'
type Definition struct {
	Services []struct {
		Repo   string `json:"repo"`
		Proto  string `json:"proto"validate:"required"`
		Commit string `json:"commit,omitempty"`
		Branch string `json:"branch,omitempty"`
		Out    []struct {
			Path     string `json:"path"validate:"required"`
			Language string `json:"language"validate:"required,language"`
		} `json:"out"validate:"required,dive"`
	} `json:"services"validate:"dive"`
	Builder string `json:"builder,omitempty"`
	Syntax  string `json:"syntax"validate:"required,syntax"`
}
