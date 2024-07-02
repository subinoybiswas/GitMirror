package analytics

import (
	
	"gitmirror/db"
	"net/http"
)

func Getinfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {


	// Specific headers
	info := db.HeaderInfo{
		UserAgent:      r.Header.Get("User-Agent"),
		Referer:        r.Header.Get("Referer"),
		Host:           r.Host,
		AcceptLanguage: r.Header.Get("Accept-Language"),
		AcceptEncoding: r.Header.Get("Accept-Encoding"),
		Cookie:         r.Header.Get("Cookie"),
		Authorization:  r.Header.Get("Authorization"),
		XForwardedFor:  r.Header.Get("X-Forwarded-For"),
		XRealIP:        r.Header.Get("X-Real-IP"),
		RemoteAddr:     r.RemoteAddr,
	}
	db.SaveHeader(info)
	next.ServeHTTP(w, r)
})

}