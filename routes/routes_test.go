package routes_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/delivercodes/bikemessenger/routes"
	"github.com/delivercodes/bikemessenger/utils"
)

var respRec *httptest.ResponseRecorder

//HealthRoute ...
func TestHealthRoute(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.HealthRoute)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `ok`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

//CheckRoute ...
func TestCheckRoute(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.CheckRoute)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

//RestartRoute restarts the docker service
func TestRestartRoute(t *testing.T) {
	req, err := http.NewRequest("POST", "/restart", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.RestartRoute)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

//KillRoute kills the docker route
func TestKillRoute(t *testing.T) {
	req, err := http.NewRequest("POST", "/kill", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.KillRoute)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetConfig(t *testing.T) {
	req, err := http.NewRequest("GET", "/config", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.Config)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestPostConfig(t *testing.T) {
	data, _ := utils.LoadConfigToJSON("../data.yml")
	r := bytes.NewReader(data)
	req, err := http.NewRequest("POST", "/config", r)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.Config)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestPostEmptyConfig(t *testing.T) {
	req, err := http.NewRequest("POST", "/config", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.Config)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}
