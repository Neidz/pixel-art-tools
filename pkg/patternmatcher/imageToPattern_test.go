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
	if !PatternsAreEqual(pattern, expectedPattern) {
		t.Errorf("Expected pattern %v, but got %v", expectedPattern, pattern)
	}

	pattern, err = ImageToPattern(img, color.RGBA{R: 200, G: 200, B: 200, A: 255}, 10)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expectedPattern = nil
	if !PatternsAreEqual(pattern, expectedPattern) {
		t.Errorf("Expected pattern %v, but got %v", expectedPattern, pattern)
	}
}
