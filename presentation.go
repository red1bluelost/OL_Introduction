package main

import (
	"github.com/red1bluelost/OL_Introduction/page"
	"log"
	"net/http"
)


func main() {
	http.HandleFunc("/view/", page.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//func main() {
//	testpage := &Page{
//		Title:  "Fuck",
//		Point1: nil,
//		Point2: nil,
//		Point3: nil,
//	}
//	testparse := []byte("(p)shit(p)fuck(p)ass(p)")
//	testpage.bodyParse(testparse)
//	fmt.Printf("1: %s, 2: %s, 3: %s",
//		testpage.Point1, testpage.Point2, testpage.Point3)
//}


