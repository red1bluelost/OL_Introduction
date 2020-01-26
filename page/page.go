package page

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/red1bluelost/OL_Introduction/link"
)

type Page struct {
	Title   string
	Point   [][]byte
	Image   *ImageHolder
	Video   string
	Heading []byte
}

var GLinker link.Linker

//handler function for all http requests
func PresentationHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/presentation/"):]
	fmt.Printf("%s\n", title)

	p, err := loadPage(title)
	if err != nil {
		log.Panicf("PresentationHandler failed %s", err)
	}
	t, _ := template.ParseFiles("data/template.html")
	t.Execute(w, p)
	//handle this error from execute
}

func StartHandler(w http.ResponseWriter, r *http.Request) {
	linker := new(link.Linker)
	linker = &GLinker
	if !linker.AlreadyInitialized {
		linker.Initialize()
	}
	linker.Reset()

	http.Redirect(w, r, "/presentation/intro", http.StatusSeeOther)
}

func IgnoreFaviconHandler(w http.ResponseWriter, r *http.Request) {
	log.Panicf("Ignore this favicon request")
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
			if p.Image == nil {
				p.Image = new(ImageHolder)
			}
			p.Image.AssignImage(rawData[:(end - 3)])
		case VIDEO:
			p.Video = string(rawData[:(end - 3)])
		case HEADING:
			p.Heading = rawData[:(end - 3)]
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

	ok = bytes.Compare(HEADING.FlagBytes(), rawData)
	if ok == 0 {
		return HEADING, true
	}

	return 0, false
}
