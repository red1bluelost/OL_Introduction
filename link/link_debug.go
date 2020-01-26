package link

import "fmt"

func (l *Linker) DebugLinks() {
	fmt.Printf("Already Initialized: %t\n", l.AlreadyInitialized)
	fmt.Printf("previous link: %s\n",l.previousLink)
	fmt.Println("")

	fmt.Printf("viewSchool: %s\n",l.viewSchool.printLink())
	fmt.Printf("school: %s\n",l.school.printLink())
	fmt.Printf("sch: %d\n", l.sch)
	fmt.Printf("schCecs: %s\n",l.schCecs.printLink())
	fmt.Printf("schNUCAR: %s\n",l.schNUCAR.printLink())
	fmt.Printf("schClubs: %s\n",l.schClubs.printLink())
	fmt.Println("")

	fmt.Printf("viewMusic: %s\n",l.viewMusic.printLink())
	fmt.Printf("music: %s\n",l.music.printLink())
	fmt.Printf("mus: %d\n", l.mus)
	fmt.Printf("musGuitar: %s\n",l.musGuitar.printLink())
	fmt.Printf("musMetal: %s\n",l.musMetal.printLink())
	fmt.Printf("musStrEdg: %s\n",l.musStrEdg.printLink())
	fmt.Println("")

	fmt.Printf("end: %s\n", l.end.printLink())
}

func (l *link) printLink() string {
	if l == nil {
		return "<nil>"
	}
	return string(*l)
}
