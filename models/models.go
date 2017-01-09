package models

import (
	"fmt"
	"os/user"
)

//Config struct is the object that comes from the yaml file
type Config struct {
	Service map[string]Service `json:"services"`
}

//Service ldfdf
type Service struct {
	Image string   `json:"image"`
	Ports []string `json:"ports"`
	Env   []string `json:"env"`
}

//ConfigFile ..fdsfd
func ConfigFile() string {
	usr, _ := user.Current()
	dir := usr.HomeDir
	return fmt.Sprintf(dir + "/.bikemessenger.yml")
}
