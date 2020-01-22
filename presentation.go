package main

import (
	"github.com/red1bluelost/OL_Introduction/page"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", page.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
