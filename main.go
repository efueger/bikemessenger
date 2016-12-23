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
	services.PullService()
	log.Fatal(http.ListenAndServe(":9000", nil))

}
