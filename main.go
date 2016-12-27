package main

import (
	"log"
	"net/http"

	"github.com/delivercodes/bikemessenger/routes"
	"github.com/delivercodes/bikemessenger/services"
)

func main() {
	http.HandleFunc("/", routes.CheckRoute)
	http.HandleFunc("/health", routes.HealthRoute)
	http.HandleFunc("/restart", routes.RestartRoute)
	services.PullService()
	log.Fatal(http.ListenAndServe(":4000", nil))

}
