package services_test

import (
	"testing"

	"github.com/delivercodes/bikemessenger/services"
)

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
