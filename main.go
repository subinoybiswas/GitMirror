package main

import (
	"gitmirror/internal"
	"gitmirror/logging"
	"log"
	"net/http"
)


func main() {
	mux := http.NewServeMux()
	serviceHandler:= http.HandlerFunc(internal.ServiceHandler)
	mux.Handle("/", logging.LoggingMiddleware(serviceHandler))
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
