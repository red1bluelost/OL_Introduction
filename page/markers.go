package page

//enumerated type for flags marking information in the text files
type Flag int

const (
	POINT Flag = 0
	IMAGE Flag = 1
	VIDEO Flag = 2
	HEADING Flag = 3
)

func (f Flag) FlagBytes() []byte {
	flags := [...][]byte{
		[]byte("(p)"),
		[]byte("(i)"),
		[]byte("(v)"),
		[]byte("(h)"),
	}
	return flags[f]
}
