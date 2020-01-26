package link

import ()

type Linker struct {
	AlreadyInitialized bool
	previousLink       string

	intro *link

	school   *link
	sch      uint8
	schCecs  *link
	schNUCAR *link
	schClubs *link

	music     *link
	mus       uint8
	musMetal  *link
	musGuitar *link
	musStrEdg *link

	end *link
}

type link string

func (l *link) activate() {
	*l = "active"
}

func (l *Linker) Initialize() {
	l.intro = new(link)

	l.school = new(link)
	l.schCecs = new(link)
	l.schNUCAR = new(link)
	l.schClubs = new(link)

	l.music = new(link)
	l.musMetal = new(link)
	l.musGuitar = new(link)
	l.musStrEdg = new(link)

	l.end = new(link)
}

func (l *Linker) Reset() {
	l.sch, l.mus = 2, 2
	l.intro.activate()

	l.school = nil
	l.schCecs = nil
	l.schNUCAR = nil
	l.schClubs = nil

	l.music = nil
	l.musMetal = nil
	l.musGuitar = nil
	l.musStrEdg = nil

	l.end = nil
}
