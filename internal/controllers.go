package internal

import (
	"fmt"
	"gitmirror/extractor"
	"log"
	"net/http"
)

func ServiceHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	username, ok := r.Context().Value(extractor.UsernameKey).(string)
    if !ok {
        http.Error(w, "Username not found in context", http.StatusInternalServerError)
        return
    }

    repo, ok := r.Context().Value(extractor.RepoKey).(string)
    if !ok {
        repo = ""
    }
	
	fmt.Println("Username:", username,"Repo:", repo)

	
	services := [2]string{"github.com", "gitlab.com"}

	for _, service := range services {
		serviceURL := "https://" + service + path
		log.Println("Checking", serviceURL,path)

		if CheckService(serviceURL) {
			log.Println("Redirecting to", serviceURL)
			http.Redirect(w, r, serviceURL, http.StatusFound)
			return
		}
	}

	fileServer := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix(path, http.FileServer(http.Dir("static"))).ServeHTTP(w, r)
	})
	fileServer.ServeHTTP(w, r)

	
}