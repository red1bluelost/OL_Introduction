package page

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Flag int

const (
	POINT Flag = 0
	IMAGE Flag = 1
	VIDEO Flag = 2
)

func (f Flag) FlagBytes() []byte {
	flags := [...][]byte{
		[]byte("(p)"),
		[]byte("(i)"),
		[]byte("(v)"),
	}
	return flags[f]
}

type Page struct {
	Title string
	Point [][]byte
}

func Handler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/"):]
	p, err := handleLoadPage(title)
	if err != nil {
		log.Panicf("Handler failed %s", err)
	}
	t, _ := template.ParseFiles("kirby.html")
	t.Execute(w, p)
}

func handleLoadPage(title string) (*Page, error) {
	title = title[:]
	page, err := loadPage(title)
	return page, err
}

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

func (p *Page) BodyParse(df []byte) error {
	for len(df) > 3 {
		end := 0
		var curFlag Flag
		for j := 0; j < len(df) ; j++ {
			var ok bool
			curFlag, ok = checkFlag(df[j:(j + 3)])
			if ok {
				end = j + 3
				break
			}
		}
		switch curFlag {
		case POINT:
			p.Point = append(p.Point, df[:(end-3)])
		case IMAGE:
			panic("implement this")
		case VIDEO:
			panic("implement this")
		}
		df = df[end:]
		fmt.Printf("%s\n", df)
	}
	return nil
}

func checkFlag(df []byte) (Flag, bool) {
	ok := bytes.Compare(POINT.FlagBytes(), df)
	if ok == 0 {
		return POINT, true
	}
	ok = bytes.Compare(IMAGE.FlagBytes(), df)
	if ok == 0 {
		return IMAGE, true
	}
	ok = bytes.Compare(VIDEO.FlagBytes(), df)
	if ok == 0 {
		return VIDEO, true
	}
	return 0, false
}
