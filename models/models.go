package models

//Config struct is the object that comes from the yaml file
type Config struct {
	Service struct {
		Image string   `json:"image"`
		Ports []string `json:",flow"`
		Env   []string `json:",flow"`
	}
}
