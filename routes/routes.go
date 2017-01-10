package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/delivercodes/bikemessenger/models"
	"github.com/delivercodes/bikemessenger/services"
	"github.com/delivercodes/bikemessenger/utils"
)

//HealthRoute ...
func HealthRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

//CheckRoute ...
func CheckRoute(w http.ResponseWriter, r *http.Request) {
	out, _ := services.CheckService()
	fmt.Fprintf(w, "%s", out)
}

//RestartRoute restarts the docker service
func RestartRoute(w http.ResponseWriter, r *http.Request) {
	services.RestartService(r.URL.Query().Get("id")).Run()
	out, _ := services.CheckService()
	fmt.Fprintf(w, "%s", out)
}

//KillRoute kills the docker route
func KillRoute(w http.ResponseWriter, r *http.Request) {
	out, _ := services.KillService(r.URL.Query().Get("id"))
	fmt.Fprintf(w, "%s", out)
}

//Config is a get request to get the config
func Config(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getConfig(w, r)
	} else if r.Method == "POST" {
		postConfig(w, r)
	}
}

func getConfig(w http.ResponseWriter, r *http.Request) {
	json, _ := utils.LoadConfigToJSON(models.ConfigFile())
	fmt.Fprintf(w, "%s", json)
}

func postConfig(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
	} else {
		var config models.Config
		json.NewDecoder(r.Body).Decode(&config)

		defer r.Body.Close()
		json, _ := utils.SaveConfigToFile(config)
		fmt.Fprintf(w, "%s", json)
	}

}
