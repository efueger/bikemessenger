package main

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

func readfile(file string) Config {
	t := Config{}
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	ymlerr := yaml.Unmarshal([]byte(dat), &t)
	if ymlerr != nil {
		log.Fatalf("error: %v", err)
	}
	return t
}
