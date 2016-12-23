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
