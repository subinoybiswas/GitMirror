package extractor

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

type contextKey string

const (
    UsernameKey contextKey = "username"
    RepoKey     contextKey = "repo"
)

func PathMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        path := r.URL.Path
        pat := strings.Split(path, "/")
        if len(pat) < 2 {
            http.Error(w, "Invalid path", http.StatusBadRequest)
            return
        }
        username := pat[1]

        var repo string
        if len(pat) >= 3 {
            repo = pat[2]
        }

        fmt.Println("Username:", username, "Repo:", repo)

        // Add extracted values to the request context
        ctx := context.WithValue(r.Context(), UsernameKey, username)
        ctx = context.WithValue(ctx, RepoKey, repo)
        
        // Call the next handler with the new context
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}