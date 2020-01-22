package page

type ImageHolder struct {
	//fill in sections as you add images
	//possibly replace numbers with more descriptive names
	Img1, Img2, Img3 *string
}

//takes input parses from txt file and assigns to specific image holder
//to signal which parts of the html template to use
func (ih *ImageHolder) AssignImage(img []byte) {
	input := string(img)

	switch input {
	//possibly replace numbers with more descriptive names
	case "Image 1":
		ih.Img1 = &input
	case "Image 2":
		ih.Img2 = &input
	case "Image 3":
		ih.Img3 = &input
	}

}
