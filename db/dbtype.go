package db

type UserMapping struct {
    ID             int
    GitHubUsername string
    GitLabUsername string

}
type HeaderInfo struct {
	UserAgent      string
	Referer        string
	Host           string
	AcceptLanguage string
	AcceptEncoding string
	Cookie         string
	Authorization  string
	XForwardedFor  string
	XRealIP        string
	RemoteAddr     string
}