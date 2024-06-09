package internal

import (
	"log"
	"net/http"
)

func CheckService(url string) bool {
   resp,err:= http.Get(url)
   log.Println("URL:", url)
   log.Println("Response:", resp.StatusCode)
   if err != nil {
	  return false
   }
   if resp.StatusCode >= 200 && resp.StatusCode < 300 {
	  return true
   }
   return false
}