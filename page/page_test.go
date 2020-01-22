package page

import (
	"bytes"
	"strings"
	"testing"
)

func TestBodyParse(t *testing.T) {
	testPage := new(Page)
	testInput := []byte(
		"Point one(p)Image 1(i)Point two(p)Image 2(i)Point three(p)Point four(p)video.link.to.a.youtube.video(v)",
	)
	resultPoint := [][]byte{[]byte("Point one"), []byte("Point two"), []byte("Point three"), []byte("Point four")}
	resultVideo := "video.link.to.a.youtube.video"

	testPage.BodyParse(testInput)

	if len(testPage.Point) != len(resultPoint) {
		t.Errorf(
			"Incorrect amount of points in the page. Should be %d, was %d",
			len(resultPoint), len(testPage.Point),
		)
	}
	for i := 0; i < len(resultPoint); i++ {
		if bytes.Compare(resultPoint[i], testPage.Point[i]) != 0 {
			t.Errorf(
				"Point %d should be %s, was %s.",
				i, resultPoint[i], testPage.Point[i],
			)
		}
	}
	if testPage.Image == nil {
		t.Errorf("Images were not found.")
	} else {
		if testPage.Image.Img1 == nil {
			t.Errorf("Image 1 was not found.")
		}
		if testPage.Image.Img2 == nil {
			t.Errorf("Image 2 was not found.")
		}
		if testPage.Image.Img3 != nil {
			t.Errorf("Image 3 was somehow found.")
		}
	}
	if strings.Compare(resultVideo, testPage.Video) != 0 {
		t.Errorf(
			"Video should be %s, was %s.",
			resultVideo, testPage.Video,
		)
	}
}
