package link

import "fmt"

func (l *Linker) DebugLinks() {
	fmt.Printf("previous link: %s\n",l.previousLink)
	fmt.Println("")

	fmt.Printf("ViewSchool: %s\n",l.ViewSchool.printLink())
	fmt.Printf("School: %s\n",l.School.printLink())
	fmt.Printf("sch: %d\n", l.sch)
	fmt.Printf("SchCecs: %s\n",l.SchCecs.printLink())
	fmt.Printf("SchNUCAR: %s\n",l.SchNUCAR.printLink())
	fmt.Printf("SchClubs: %s\n",l.SchClubs.printLink())
	fmt.Println("")

	fmt.Printf("ViewMusic: %s\n",l.ViewMusic.printLink())
	fmt.Printf("Music: %s\n",l.Music.printLink())
	fmt.Printf("mus: %d\n", l.mus)
	fmt.Printf("MusGuitar: %s\n",l.MusGuitar.printLink())
	fmt.Printf("MusMetal: %s\n",l.MusMetal.printLink())
	fmt.Printf("MusStrEdg: %s\n",l.MusStrEdg.printLink())
	fmt.Println("")

	fmt.Printf("End: %s\n", l.End.printLink())
}

func (l *link) printLink() string {
	if l == nil {
		return "<nil>"
	}
	return string(*l)
}
