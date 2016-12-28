package routes

import (
	"fmt"
	"net/http"

	"github.com/delivercodes/bikemessenger/services"
)

//HealthRoute ...
func HealthRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

//CheckRoute ...
func CheckRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", services.CheckService())

}

//RestartRoute restarts the docker service
func RestartRoute(w http.ResponseWriter, r *http.Request) {
	out := services.RestartService(r.URL.Query().Get("id"))
	fmt.Fprintf(w, "%s", out)
}

//KillRoute kills the docker route
func KillRoute(w http.ResponseWriter, r *http.Request) {
	out, err := services.KillService(r.URL.Query().Get("id"))
	if err != nil {

	}
	fmt.Fprintf(w, "%s", out)
}
