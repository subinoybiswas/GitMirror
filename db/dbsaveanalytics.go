package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/subinoybiswas/goenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)
func SaveHeaderInfo(db *sql.DB, info HeaderInfo) error {
	query := `
		INSERT INTO headers (user_agent, referer, host, accept_language, accept_encoding, cookie, authorization, x_forwarded_for, x_real_ip, remote_addr)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := db.Exec(query, info.UserAgent, info.Referer, info.Host, info.AcceptLanguage, info.AcceptEncoding, info.Cookie, info.Authorization, info.XForwardedFor, info.XRealIP, info.RemoteAddr)
	if err != nil {
		return fmt.Errorf("failed to insert header info: %v", err)
	}
	return nil
}


func SaveHeader(info HeaderInfo) {
	value, err := goenv.GetEnv("TORSO_STRING")
	
	if err != nil {
		fmt.Println(err)
	}
  
	url :=value
  
	db, err := sql.Open("libsql", url)
	if err != nil {
	  fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)

	}
	
	err1 := SaveHeaderInfo(db, info)
	
	if err1 != nil {
	  fmt.Fprintf(os.Stderr, "%s", err1)
	  
	}
	defer db.Close()
  

  }