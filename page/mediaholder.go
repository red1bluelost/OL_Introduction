package page

type ImageHolder struct {
	//fill in sections as you add images
	LV, AVHS *string
	CPP, GitHub, GoLang, Akita, Baja, NUSound *string
	RG, BTBAM, AAL, Paul, Alissa *string
	Program *string
}

//takes input parses from txt file and assigns to specific image holder
//to signal which parts of the html template to use
func (ih *ImageHolder) AssignImage(img []byte) {
	input := string(img)

	switch input {
	case "LasVegas":
		ih.LV = &input
	case "ArborView":
		ih.AVHS = &input
	case "CPP":
		ih.CPP = &input
	case "GitHub":
		ih.GitHub = &input
	case "GoLang":
		ih.GoLang = &input
	case "Akita":
		ih.Akita = &input
	case "Baja":
		ih.Baja = &input
	case "NUSound":
		ih.NUSound = &input
	case "RG":
		ih.RG = &input
	case "BTBAM":
		ih.BTBAM = &input
	case "AAL":
		ih.AAL = &input
	case "Paul":
		ih.Paul = &input
	case "Alissa":
		ih.Alissa = &input
	case "meme":
		ih.Program = &input
	}

}
