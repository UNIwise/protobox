package yaml

// Definition describe the configuration yaml file 'protobuf.yaml'
type Definition struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}
