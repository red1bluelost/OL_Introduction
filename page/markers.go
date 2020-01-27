package page

//enumerated type for flags marking information in the text files
type Flag int

const (
	POINT   Flag = 0
	IMAGE   Flag = 1
	HEADING Flag = 2
)

func (f Flag) FlagBytes() []byte {
	flags := [...][]byte{
		[]byte("(p)"),
		[]byte("(i)"),
		[]byte("(h)"),
	}
	return flags[f]
}
