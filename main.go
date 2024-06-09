package main

import (
	"gitmirror/internal"
	"log"
	"net/http"
)



func main() {
    http.HandleFunc("/", internal.ServiceHandler)
    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
