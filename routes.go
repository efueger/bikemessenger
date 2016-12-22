package main

import (
	"fmt"
	"net/http"
)

func healthRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

func checkRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", checkService())

}
