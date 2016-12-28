package routes

import (
	"fmt"
	"io"
	"net/http"

	"github.com/delivercodes/bikemessenger/services"
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
	fw := flushWriter{w: w}
	if f, ok := w.(http.Flusher); ok {
		fw.f = f
	}
	cmd := services.RestartService(r.URL.Query().Get("id"))
	cmd.Stdout = &fw
	cmd.Stderr = &fw
	cmd.Run()
}

//KillRoute kills the docker route
func KillRoute(w http.ResponseWriter, r *http.Request) {
	out, err := services.KillService(r.URL.Query().Get("id"))
	if err != nil {

	}
	fmt.Fprintf(w, "%s", out)
}
