package utils_test

import (
	"testing"

	"github.com/delivercodes/bikemessenger/models"
	"github.com/delivercodes/bikemessenger/utils"
)

func TestLoadConfigToModel(t *testing.T) {
	config, err := utils.LoadConfigToModel(models.ConfigFile())
	if err != nil {
		t.Error("Error Config didn't load ", err)
	}
	for service := range config.Service {
		if config.Service[service].Image == "" {
			t.Error("Error Config didn't load ", service)
		}
	}
}

func TestLoadConfigToJSON(t *testing.T) {
	j, err := utils.LoadConfigToJSON()
	if err != nil {
		t.Error("Error JSON didn't load ", j)
	}
	if j == nil {
		t.Error("Error JSON didn't load ", j)
	}
}

func TestSaveConfigToFile(t *testing.T) {
	config, _ := utils.LoadConfigToModel(models.ConfigFile())
	j, err := utils.SaveConfigToFile(config)
	if err != nil {
		t.Error("Error File didn't save ", j)
	}
}
