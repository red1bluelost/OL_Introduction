package link

import "fmt"

type Linker struct {
	//AlreadyInitialized bool
	previousLink       string

	ViewSchool *link
	School     *link
	sch        uint8
	SchCecs    *link
	SchNUCAR   *link
	SchClubs   *link

	ViewMusic *link
	Music     *link
	mus       uint8
	MusGuitar *link
	MusMetal  *link
	MusStrEdg *link

	End *link
}

type link string

var gLink link = "active"

func (l *Linker) Reset() {
	l.sch, l.mus = 4, 4

	l.ViewSchool = nil
	l.School = nil
	l.SchCecs = nil
	l.SchNUCAR = nil
	l.SchClubs = nil

	l.ViewMusic = nil
	l.Music = nil
	l.MusMetal = nil
	l.MusGuitar = nil
	l.MusStrEdg = nil

	l.End = nil
}

func (l *Linker) Handle(s string) {
	if s == l.previousLink {
		return
	}
	l.previousLink = s

	switch s {
	case "intro":
		l.Music = &gLink
		l.School = &gLink
		l.ViewMusic = &gLink
		l.ViewSchool = &gLink
		return
	case "school", "CECS", "NUCAR", "clubs":
		l.handleSchool(s)
		return
	case "music", "guitar", "metal", "stredg":
		l.handleMusic(s)
		return
	case "end":
		l.Reset()
		return
	default:
		fmt.Printf("Linker does not know how to handle %s\n", s)
	}
}

func (l *Linker) handleSchool(s string) {
	l.sch--

	if l.sch == 0 && l.mus == 0 {
		l.setEnd()
		return
	}

	switch s {
	case "school":
		l.School = nil
		l.SchCecs = &gLink
		l.SchClubs = &gLink
		l.SchNUCAR = &gLink
		l.ViewMusic = nil
		return
	case "CECS":
		l.SchCecs = nil
		if l.sch == 0 {l.ViewMusic, l.ViewSchool = &gLink, nil}
		return
	case "clubs":
		l.SchClubs = nil
		if l.sch == 0 {l.ViewMusic, l.ViewSchool = &gLink, nil}
		return
	case "NUCAR":
		l.SchNUCAR = nil
		if l.sch == 0 {l.ViewMusic, l.ViewSchool = &gLink, nil}
		return
	}
}

func (l *Linker) handleMusic(s string) {
	l.mus--

	if l.sch == 0 && l.mus == 0 {
		l.setEnd()
		return
	}

	switch s {
	case "music":
		l.Music = nil
		l.MusGuitar = &gLink
		l.ViewSchool = nil
		return
	case "guitar":
		l.MusGuitar = nil
		l.MusMetal = &gLink
		return
	case "metal":
		l.MusMetal = nil
		l.MusStrEdg = &gLink
		return
	case "stredg":
		l.MusStrEdg = nil
		l.ViewSchool = &gLink
		l.ViewMusic = nil
		return
	}
}

func (l *Linker) setEnd() {
	l.Reset()
	l.End = &gLink
}
