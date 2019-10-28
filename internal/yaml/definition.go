package yaml

// Definition describe the configuration yaml file 'protobuf.yaml'
type Definition struct {
	Services []struct {
		Repo   string `json:"repo"`
		Proto  string `json:"proto"`
		Commit string `json:"commit,omitempty,default origin/HEAD"`
		Branch string `json:"branch,omitempty,default master"`
		Out    []struct {
			Path     string `json:"path"`
			Language string `json:"language"`
		} `json:"out"`
	} `json:"services"`
	Builder string `json:"builder,omitempty"`
	Syntax  string `json:"syntax"`
}
