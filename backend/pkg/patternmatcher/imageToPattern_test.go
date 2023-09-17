package patternmatcher

import (
	"image"
	"image/color"
	"testing"
)

func TestImageToPattern(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 3, 3))
	img.Set(1, 1, color.RGBA{R: 255, G: 255, B: 255, A: 255})

	pattern, err := ImageToPattern(img, color.RGBA{R: 255, G: 255, B: 255, A: 255}, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expectedPattern := Pattern{{1, 1}}
	if !patternsAreEqual(pattern, expectedPattern) {
		t.Errorf("Expected pattern %v, but got %v", expectedPattern, pattern)
	}

	pattern, err = ImageToPattern(img, color.RGBA{R: 200, G: 200, B: 200, A: 255}, 10)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expectedPattern = nil
	if !patternsAreEqual(pattern, expectedPattern) {
		t.Errorf("Expected pattern %v, but got %v", expectedPattern, pattern)
	}
}

func patternsAreEqual(p1, p2 Pattern) bool {
	if len(p1) != len(p2) {
		return false
	}
	for i := range p1 {
		if p1[i] != p2[i] {
			return false
		}
	}
	return true
}
