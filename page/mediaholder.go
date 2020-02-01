package page

import (
	"bytes"
	"fmt"
)

type ImageHolder struct {
	//fill in sections as you add images
	LV, AVHS bool
	CPP, GitHub, GoLang, Akita, Baja, NUSound bool
	RG, BTBAM, AAL, Paul, Alissa bool
	Program bool
}


//takes input parses from txt file and assigns to specific image holder
//to signal which parts of the html template to use
func (ih *ImageHolder) AssignImage(img []byte) {
	switch {
	case 0 == bytes.Compare([]byte("LasVegas"), img):
		fmt.Printf("Found Vegas\n")
		ih.LV = true
	case 0 == bytes.Compare([]byte("ArborView"), img):
		ih.AVHS = true
	case 0 == bytes.Compare([]byte("CPP"), img):
		ih.CPP = true
	case 0 == bytes.Compare([]byte("GitHub"), img):
		ih.GitHub = true
	case 0 == bytes.Compare([]byte("GoLang"), img):
		ih.GoLang = true
	case 0 == bytes.Compare([]byte("Akita"), img):
		ih.Akita = true
	case 0 == bytes.Compare([]byte("Baja"), img):
		ih.Baja = true
	case 0 == bytes.Compare([]byte("NUSound"), img):
		ih.NUSound = true
	case 0 == bytes.Compare([]byte("RG"), img):
		ih.RG = true
	case 0 == bytes.Compare([]byte("BTBAM"), img):
		ih.BTBAM = true
	case 0 == bytes.Compare([]byte("AAL"), img):
		ih.AAL = true
	case 0 == bytes.Compare([]byte("Paul"), img):
		ih.Paul = true
	case 0 == bytes.Compare([]byte("Alissa"), img):
		ih.Alissa = true
	case 0 == bytes.Compare([]byte("meme"), img):
		ih.Program = true
	default:
		fmt.Printf("Image handler could not assign: %s\n", img)
	}
}
