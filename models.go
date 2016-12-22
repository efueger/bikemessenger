package main

//Config ...
type Config struct {
	Service struct {
		Image string   `yaml:"image"`
		Ports []string `yaml:",flow"`
		Env   []string `yaml:",flow"`
	}
}
