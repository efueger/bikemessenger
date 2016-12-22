package main

//Config ...
type Config struct {
	Service struct {
		Image string   `yaml:"image"`
		Ports []string `yaml:",flow"`
		Env   []string `yaml:",flow"`
	}
}

//
// //Config ...
// type Config struct {
// 	A string
// 	B struct {
// 		RenamedC int   `yaml:"c"`
// 		D        []int `yaml:",flow"`
// 	}
// }
