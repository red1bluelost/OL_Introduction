package link

import ()

type Linker struct {
	intro *link

	college  *link
	col      uint8
	colCecs  *link
	colNUCAR *link
	colClubs *link

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

var G_Linker Linker

func (l *Linker) Initialize() {
	l.intro = new(link)

	l.college = new(link)
	l.colCecs = new(link)
	l.colNUCAR = new(link)
	l.colClubs = new(link)

	l.music = new(link)
	l.musMetal = new(link)
	l.musGuitar = new(link)
	l.musStrEdg = new(link)

	l.end = new(link)
}

func (l *Linker) Reset() {
	l.col, l.mus = 2, 2
	l.intro.activate()

	l.college = nil
	l.colCecs = nil
	l.colNUCAR = nil
	l.colClubs = nil

	l.music = nil
	l.musMetal = nil
	l.musGuitar = nil
	l.musStrEdg = nil

	l.end = nil
}
