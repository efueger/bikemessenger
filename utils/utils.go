package utils

import (
	"io/ioutil"
	"log"

	"github.com/delivercodes/bikemessenger/models"
	yaml "gopkg.in/yaml.v2"
)

//Readfile ...
func Readfile(file string) models.Config {
	t := models.Config{}
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	ymlerr := yaml.Unmarshal([]byte(dat), &t)
	if ymlerr != nil {
		log.Fatalf("error loading yaml to model: %v", err)
	}
	return t
}
