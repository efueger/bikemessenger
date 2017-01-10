package services_test

import (
	"testing"

	"github.com/delivercodes/bikemessenger/services"
	"github.com/delivercodes/bikemessenger/utils"
)

func TestPullService(t *testing.T) {
	config, _ := utils.LoadConfigToModel("../data.yml")
	services.PullService(config)
}

func TestCheckService(t *testing.T) {
	services.CheckService()
}

func TestRestartService(t *testing.T) {
	cmd := services.RestartService("alpine")
	cmd.Run()
}
