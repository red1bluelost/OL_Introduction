package page

import (
	"github.com/red1bluelost/OL_Introduction/image"
	"html/template"
	"log"
	"net/http"
	"reflect"
)

type Pager interface {
	BodyParse(df []byte) error
}

func Handler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := handleLoadPage(title)
	if err != nil {
		log.Panicf("Handler failed %s", err)
	}
	switch p := p.(type) {
	case *image.ImagePage:
		t, _ := template.ParseFiles("kirby.html")
		t.Execute(w, p)
	default:
		log.Panicf("cannot handle page of type (%s)", reflect.TypeOf(p))
	}
}

func handleLoadPage(title string) (Pager, error) {
	switch title[:5] {
	case "image":
		title = title[6:]
		page, err := image.LoadImagePage(title)
		return page, err
	default:
		panic("handleLoadPage failed")
	}
	return nil, nil
}

