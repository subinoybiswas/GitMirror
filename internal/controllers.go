package internal

import (
	"log"
	"net/http"
)


func ServiceHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	services := [2]string{"github.com", "gitlab.com"}

	for _, service := range services {
		serviceURL := "https://" + service + path
		log.Println("Checking", serviceURL)
		if CheckService(serviceURL) {
			log.Println("Redirecting to", serviceURL)
			http.Redirect(w, r, serviceURL, http.StatusFound)
			return
		}

	
	}
	http.NotFound(w, r)
}