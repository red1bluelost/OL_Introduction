package main

import (
	"log"
	"net/http"

	"github.com/red1bluelost/OL_Introduction/page"
)

func main() {
	//redirect css and images requests
	http.Handle(
		"/presentation/css/",
		http.StripPrefix("/presentation/css/", http.FileServer(http.Dir("data/css"))),
	)
	http.Handle(
		"/css/",
		http.StripPrefix("/css/", http.FileServer(http.Dir("data/css"))),
	)
	http.Handle(
		"/presentation/images/",
		http.StripPrefix("/presentation/images/", http.FileServer(http.Dir("data/images"))),
	)

	//handle urls
	http.HandleFunc("/start", page.StartHandler)
	http.HandleFunc("/presentation/", page.PresentationHandler)
	http.HandleFunc("/links", page.LinkPageHandler)


	//http.HandleFunc("favicon", page.IgnoreFaviconHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
