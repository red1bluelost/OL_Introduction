package main

import (
	"github.com/red1bluelost/OL_Introduction/page"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/start", page.StartHandler)
	http.HandleFunc("/", page.MainHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
