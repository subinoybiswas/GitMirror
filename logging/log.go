package logging

import (
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Executing Logging")
		next.ServeHTTP(w, r)
		log.Print("Executing Logging again")
		
	})
}









































