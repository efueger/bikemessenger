package models

//Config struct is the object that comes from the yaml file
type Config struct {
	Service map[string]Service `json:"services"`
}

//ConfigFile ..fdsfd
var ConfigFile string

//Service ldfdf
type Service struct {
	Image string   `json:"image"`
	Ports []string `json:"ports"`
	Env   []string `json:"env"`
}
