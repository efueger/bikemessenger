package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", checkRoute)
	http.HandleFunc("/health", healthRoute)
	pullService()
	log.Fatal(http.ListenAndServe(":9000", nil))

}
