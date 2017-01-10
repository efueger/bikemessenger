package services_test

import (
	"testing"

	"github.com/delivercodes/bikemessenger/models"
	"github.com/delivercodes/bikemessenger/services"
	"github.com/delivercodes/bikemessenger/utils"
)

func TestRunService(t *testing.T) {
	config, _ := utils.LoadConfigToModel(models.ConfigFile())
	service := config.Service["postgres"]
	cmd := services.RunService(service, "postgres")

	cmd.Start()
}

func TestPullService(t *testing.T) {
	services.PullService()
}

func TestCheckService(t *testing.T) {
	services.CheckService()
}

func TestRestartService(t *testing.T) {
	_, err := services.RestartService("postgres")
	if err != nil {
		t.Fatal("Error restarting service")
	}
}
