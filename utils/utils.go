package utils

import (
	"io/ioutil"

	"github.com/delivercodes/bikemessenger/models"
	"github.com/ghodss/yaml"
)

//LoadConfigToModel loads config and returns model
func LoadConfigToModel(file string) (models.Config, error) {
	t := models.Config{}
	dat, err := ioutil.ReadFile(file)
	yaml.Unmarshal([]byte(dat), &t)

	return t, err
}

//LoadConfigToJSON gets the config model and returns json
func LoadConfigToJSON() ([]byte, error) {
	config, err := LoadConfigToModel(models.ConfigFile())

	y, _ := yaml.Marshal(config)

	j2, _ := yaml.YAMLToJSON(y)
	return j2, err
}

//SaveConfigToFile takes Config and saves it to the data.yml file
func SaveConfigToFile(config models.Config) ([]byte, error) {
	y, err := yaml.Marshal(config)

	ioutil.WriteFile(models.ConfigFile(), y, 0644)
	j, _ := yaml.YAMLToJSON(y)
	return j, err
}
