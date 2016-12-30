package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/delivercodes/bikemessenger/models"
	"github.com/delivercodes/bikemessenger/services"
	"github.com/delivercodes/bikemessenger/utils"
)

type flushWriter struct {
	f http.Flusher
	w io.Writer
}

func (fw *flushWriter) Write(p []byte) (n int, err error) {
	n, err = fw.w.Write(p)
	if fw.f != nil {
		fw.f.Flush()
	}
	return
}

//HealthRoute ...
func HealthRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

//CheckRoute ...
func CheckRoute(w http.ResponseWriter, r *http.Request) {
	out, err := services.CheckService()
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}
	fmt.Fprintf(w, "%s", out)

}

//RestartRoute restarts the docker service
func RestartRoute(w http.ResponseWriter, r *http.Request) {

	services.RestartService(r.URL.Query().Get("id")).Run()

	out, err := services.CheckService()
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}
	fmt.Fprintf(w, "%s", out)
}

//KillRoute kills the docker route
func KillRoute(w http.ResponseWriter, r *http.Request) {
	out, err := services.KillService(r.URL.Query().Get("id"))
	if err != nil {

	}
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
	json := utils.LoadConfigToJSON()
	fmt.Fprintf(w, "%s", json)
}

func postConfig(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var config models.Config
	err := decoder.Decode(&config)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	fmt.Fprintf(w, "%s", config)
}
