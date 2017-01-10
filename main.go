package main

import (
	"log"
	"net/http"
	"time"

	"github.com/delivercodes/bikemessenger/models"
	"github.com/delivercodes/bikemessenger/routes"
	"github.com/delivercodes/bikemessenger/services"
	"github.com/delivercodes/bikemessenger/utils"
	"github.com/gorilla/mux"
)

//Router dfdfd
func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", routes.CheckRoute)
	r.HandleFunc("/health", routes.HealthRoute)

	//Status
	r.HandleFunc("/restart", routes.RestartRoute)
	r.HandleFunc("/kill", routes.KillRoute)

	//Config
	r.HandleFunc("/config", routes.Config)
	return r
}

//Server setup for server
func Server(r *mux.Router) *http.Server {
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:4000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return srv
}

//Setup sets server up for logfatal
func Setup(config models.Config) *http.Server {
	r := Router()

	http.Handle("/", r)
	services.PullService(config)
	srv := Server(r)
	return srv
}

func main() {
	config, _ := utils.LoadConfigToModel(models.ConfigFile())
	log.Fatal(Setup(config).ListenAndServe())
}
