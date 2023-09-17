// patternsInImage_test.go

package patternmatcher

import (
	"image"
	"image/color"
	"testing"
)

func TestPatternsInImage(t *testing.T) {
	img := createSampleImage()

	pattern := Pattern{
		{0, 0},
		{1, 0},
		{2, 0},
		{0, 1},
		{2, 1},
		{0, 2},
		{1, 2},
		{2, 2},
	}

	patterns := PatternsInImage(img, pattern)
	if len(patterns) != 1 {
		t.Errorf("Expected 1 pattern, but got %d", len(patterns))
	}
	if !PatternsAreEqual(patterns[0], pattern) {
		t.Errorf("Expected pattern %v, but got %v", pattern, patterns[0])
	}

	nonExistentPattern := Pattern{
		{10, 10},
		{11, 10},
		{12, 10},
		{10, 11},
		{12, 11},
		{10, 12},
		{11, 12},
		{12, 12},
	}
	patterns = PatternsInImage(img, nonExistentPattern)
	if len(patterns) != 0 {
		t.Errorf("Expected 0 patterns, but got %d", len(patterns))
	}
}

func createSampleImage() image.Image {
	img := image.NewRGBA(image.Rect(0, 0, 3, 3))
	img.Set(0, 0, color.RGBA{R: 255, G: 255, B: 255, A: 255})
	img.Set(1, 0, color.RGBA{R: 255, G: 255, B: 255, A: 255})
	img.Set(2, 0, color.RGBA{R: 255, G: 255, B: 255, A: 255})
	img.Set(0, 1, color.RGBA{R: 255, G: 255, B: 255, A: 255})
	img.Set(2, 1, color.RGBA{R: 255, G: 255, B: 255, A: 255})
	img.Set(0, 2, color.RGBA{R: 255, G: 255, B: 255, A: 255})
	img.Set(1, 2, color.RGBA{R: 255, G: 255, B: 255, A: 255})
	img.Set(2, 2, color.RGBA{R: 255, G: 255, B: 255, A: 255})
	return img
}
