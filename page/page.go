package page

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Point [][]byte
	Image [][]byte
	Video []byte
}

//handler function for all http requests
func Handler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/"):]
	p, err := loadPage(title)
	if err != nil {
		log.Panicf("Handler failed %s", err)
	}
	t, _ := template.ParseFiles("kirby.html")
	t.Execute(w, p)
}

//takes end of url and loads page with data
func loadPage(title string) (*Page, error) {
	filename := "data/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	page := new(Page)
	ok := page.BodyParse(body)
	if ok != nil {
		panic(ok)
	}
	page.Title = title
	return page, nil
}

//parses the txt file and assigns data to correct fields
func (p *Page) BodyParse(rawData []byte) error {
	for len(rawData) > 3 {
		var (
			end     = 0
			curFlag Flag
			ok      bool
		)
		for j := 0; j < len(rawData); j++ {
			curFlag, ok = checkFlag(rawData[j:(j + 3)])
			if ok {
				end = j + 3
				break
			}
		}
		if !ok {
			return fmt.Errorf("BodyParse could not find a flag marker")
		}
		switch curFlag {
		case POINT:
			p.Point = append(p.Point, rawData[:(end-3)])
		case IMAGE:
			p.Image = append(p.Image, rawData[:(end-3)])
		case VIDEO:
			p.Video = rawData[:(end-3)]
		}
		rawData = rawData[end:]
	}
	return nil
}

//compares subsection of txt file to the flags and returns if and what type flag found
func checkFlag(rawData []byte) (Flag, bool) {
	ok := bytes.Compare(POINT.FlagBytes(), rawData)
	if ok == 0 {
		return POINT, true
	}
	ok = bytes.Compare(IMAGE.FlagBytes(), rawData)
	if ok == 0 {
		return IMAGE, true
	}
	ok = bytes.Compare(VIDEO.FlagBytes(), rawData)
	if ok == 0 {
		return VIDEO, true
	}
	return 0, false
}
