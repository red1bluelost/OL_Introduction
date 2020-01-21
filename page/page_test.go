package page

import (
	"bytes"
	"testing"
)

func TestBodyParse(t *testing.T) {
	testPage := new(Page)
	testInput := []byte("Point one(p)Point two(p)Point three(p)Point four(p)")
	testResult := [][]byte{[]byte("Point one"), []byte("Point two"), []byte("Point three"), []byte("Point four")}

	testPage.BodyParse(testInput)

	if len(testPage.Point) != len(testResult) {
		t.Errorf(
			"Incorrect amount of points in the page. Should be %d, was %d",
			len(testResult), len(testPage.Point),
		)
	}
	for i := 0; i < len(testResult); i++ {
		if bytes.Compare(testResult[i], testPage.Point[i]) != 0 {
			t.Errorf(
				"Point %d should be %s, was %s.",
				i, testResult[i], testPage.Point[i],
			)
		}
	}
}
