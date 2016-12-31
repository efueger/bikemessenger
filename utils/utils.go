package utils

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/delivercodes/bikemessenger/models"
	"github.com/ghodss/yaml"
)

//LoadConfigToModel loads config and returns model
func LoadConfigToModel(file string) models.Config {
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

//LoadConfigToJSON gets the config model and returns json
func LoadConfigToJSON() []byte {
	config := LoadConfigToModel("data.yml")
	y, err := yaml.Marshal(config)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic(err)
	}
	j2, err := yaml.YAMLToJSON(y)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic(err)
	}
	return j2
}

//SaveConfigToFile takes Config and saves it to the data.yml file
func SaveConfigToFile(config models.Config) []byte {
	y, err := yaml.Marshal(config)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil
	}
	errs := ioutil.WriteFile("data.yml", y, 0644)
	if errs != nil {
		return nil
	}
	j, err := yaml.YAMLToJSON(y)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic(err)
	}
	return j
}
