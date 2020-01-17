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

//func main() {
//	testpage := &page.Page{Title: "Fuck"}
//	testpage.Point = make([][]byte, 0)
//	testparse := []byte("shit(p)fuck(p)ass(p)")
//	testpage.BodyParse(testparse)
//
//	for i := 0; i < 3; i++ {
//		fmt.Printf("%d: %s\n", i, testpage.Point[i])
//	}
//}
