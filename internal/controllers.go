package internal

import (
	"fmt"
	"gitmirror/db"
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

	userMappings:=db.Lookup(username)

	mapping:=userMappings[0]
	
	services := [2]string{"github.com", "gitlab.com"}


	for _, service := range services {
		serviceURL := "https://" + service + "/" +  GetUsername(service,mapping) + "/" + repo
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

func GetUsername(service string,mapping db.UserMapping) (string){
	if service == "github.com" {
		return mapping.GitHubUsername
	}
	if service == "gitlab.com" {
		return mapping.GitLabUsername
	}
	return ""
}