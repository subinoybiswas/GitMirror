package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/subinoybiswas/goenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func Lookup(username string) ([]UserMapping){
  value, err := goenv.GetEnv("TORSO_STRING")
  
  if err != nil {
      fmt.Println(err)
  }

  url :=value

  db, err := sql.Open("libsql", url)
  if err != nil {
    fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
  }
  
  userMappings,err:=FindService(db, username)
  if err != nil {
    fmt.Fprintf(os.Stderr, "failed to find service for %s: %s", username, err)
    
  }
  defer db.Close()

  return userMappings
}
  



