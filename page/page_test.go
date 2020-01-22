package page

import (
	"bytes"
	"testing"
)

func TestBodyParse(t *testing.T) {
	testPage := new(Page)
	testInput := []byte(
		"Point one(p)Image_1.png(i)Point two(p)Image_2.png(i)Point three(p)Point four(p)Image_3.png(i)video.link.to.a.youtube.video(v)",
		)
	resultPoint := [][]byte{[]byte("Point one"), []byte("Point two"), []byte("Point three"), []byte("Point four")}
	resultImage := [][]byte{[]byte("Image_1.png"), []byte("Image_2.png"), []byte("Image_3.png")}
	resultVideo := []byte("video.link.to.a.youtube.video")

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
	if len(testPage.Image) != len(resultImage) {
		t.Errorf(
			"Incorrect amount of images in the page. Should be %d, was %d",
			len(resultImage), len(testPage.Image),
		)
	}
	for i := 0; i < len(resultImage); i++ {
		if bytes.Compare(resultImage[i], testPage.Image[i]) != 0 {
			t.Errorf(
				"Image %d should be %s, was %s.",
				i, resultImage[i], testPage.Image[i],
			)
		}
	}
	if bytes.Compare(resultVideo, testPage.Video) != 0 {
		t.Errorf(
			"Video should be %s, was %s.",
			resultVideo, testPage.Video,
		)
	}
}
