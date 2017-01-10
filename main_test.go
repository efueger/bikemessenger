package main

import (
	"testing"

	"github.com/delivercodes/bikemessenger/utils"
)

func TestSetup(t *testing.T) {
	config, _ := utils.LoadConfigToModel("data.yaml")
	srv := Setup(config)
	if srv == nil {
		t.Fail()
	}
}

// func TestMain(m *testing.M) {
// 	m.Run()
// }
