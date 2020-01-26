package link

type Linker struct {
	AlreadyInitialized bool
	previousLink       string

	viewSchool *link
	school     *link
	sch        uint8
	schCecs    *link
	schNUCAR   *link
	schClubs   *link

	viewMusic *link
	music     *link
	mus       uint8
	musGuitar *link
	musMetal  *link
	musStrEdg *link

	end *link
}

type link string

var gLink link = "active"

func (l *Linker) Initialize() {
	l.AlreadyInitialized = true

	l.viewSchool = new(link)
	l.school = new(link)
	l.schCecs = new(link)
	l.schNUCAR = new(link)
	l.schClubs = new(link)

	l.viewMusic = new(link)
	l.music = new(link)
	l.musMetal = new(link)
	l.musGuitar = new(link)
	l.musStrEdg = new(link)

	l.end = new(link)
}

func (l *Linker) Reset() {
	l.sch, l.mus = 4, 4

	l.viewSchool = nil
	l.school = nil
	l.schCecs = nil
	l.schNUCAR = nil
	l.schClubs = nil

	l.viewMusic = nil
	l.music = nil
	l.musMetal = nil
	l.musGuitar = nil
	l.musStrEdg = nil

	l.end = nil
}

func (l *Linker) Handle(s string) {
	if s == l.previousLink {
		return
	}
	l.previousLink = s

	switch s {
	case "intro":
		l.music = &gLink
		l.school = &gLink
		l.viewMusic = &gLink
		l.viewSchool = &gLink
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
		l.school = nil
		l.schCecs = &gLink
		l.schClubs = &gLink
		l.schNUCAR = &gLink
		l.viewMusic = nil
		return
	case "CECS":
		l.schCecs = nil
		if l.sch == 0 {l.viewMusic, l.viewSchool = &gLink, nil}
		return
	case "clubs":
		l.schClubs = nil
		if l.sch == 0 {l.viewMusic, l.viewSchool = &gLink, nil}
		return
	case "NUCAR":
		l.schNUCAR = nil
		if l.sch == 0 {l.viewMusic, l.viewSchool = &gLink, nil}
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
		l.music = nil
		l.musGuitar = &gLink
		l.viewSchool = nil
		return
	case "guitar":
		l.musGuitar = nil
		l.musMetal = &gLink
		return
	case "metal":
		l.musMetal = nil
		l.musStrEdg = &gLink
		return
	case "stredg":
		l.musStrEdg = nil
		l.viewSchool = &gLink
		l.viewMusic = nil
		return
	}
}

func (l *Linker) setEnd() {
	l.Reset()
	l.end = &gLink
}
