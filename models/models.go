package models

//Config struct is the object that comes from the yaml file
type Config struct {
	Service struct {
		Image string   `yaml:"image"`
		Ports []string `yaml:",flow"`
		Env   []string `yaml:",flow"`
	}
}
