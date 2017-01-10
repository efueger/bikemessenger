package models_test

import (
	"encoding/json"
	"testing"

	"github.com/delivercodes/bikemessenger/models"
)

func TestConfig(t *testing.T) {
	config := models.Config{}
	input := []byte(`{"services": {"postgres": {"env": ["DEBUG=2"],"image": "postgres","ports": ["8660:80","7443:443"]}}}`)
	err := json.Unmarshal(input, &config)
	if err != nil {
		return
	}
	service := config.Service["postgres"]
	testEnv := service.Env
	testPorts := service.Ports
	testImage := service.Image

	if testImage != "postgres" {
		t.Errorf("Config Model Image Error", testImage)
	}
	examplePorts := []string{"8660:80", "7443:443"}
	for i := range examplePorts {
		if examplePorts[i] != testPorts[i] {
			t.Errorf("Config Model Ports Error", testPorts)
		}
	}

	exampleEnvs := []string{"DEBUG=2"}
	for i := range exampleEnvs {
		if exampleEnvs[i] != testEnv[i] {
			t.Errorf("Config Model Env Error", testEnv)
		}
	}
}

func TestConfigFile(t *testing.T) {
	config := models.ConfigFile()
	if config != "/Users/danielcherubini/.bikemessenger.yml" {
		t.Errorf("Config Model File Location Error", config)
	}
}
