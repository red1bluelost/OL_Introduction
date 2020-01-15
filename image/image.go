package image

import (
	"bytes"
	"io/ioutil"
)

type ImagePage struct {
	Title string
	Point1 []byte
	Point2 []byte
	Point3 []byte
}

func (p *ImagePage) BodyParse(df []byte) error {
	watch := []byte("(p)")
	dflen := len(watch)
	for i := 0; i < 3; i++ {
		end := 0
		dfparse := df[dflen:]
		for j := 0; j < len(df); j++ {
			ress := bytes.Compare(watch, dfparse[j:(j+dflen)])
			if ress == 0 {
				end = j + dflen
				break
			}
		}
		switch i {
		case 0:
			p.Point1 = df[dflen:end]
		case 1:
			p.Point2 = df[dflen:end]
		case 2:
			p.Point3 = df[dflen:end]
		}
		df = df[end:]
	} //end for
	return nil
}


func LoadImagePage(title string) (*ImagePage, error) {
	filename := "data/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	page := new(ImagePage)
	ok := page.BodyParse(body)
	if ok != nil {
		panic(ok)
	}
	page.Title = title
	return page, nil
}