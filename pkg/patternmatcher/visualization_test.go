package patternmatcher

import (
	"image/color"
	"testing"
)

func TestPaintPatternsOnBlankImage(t *testing.T) {
	imgWidth := 5
	imgHeight := 5
	bgColor := color.Black
	img := filledImage(imgWidth, imgHeight, bgColor)

	pattern := Pattern{
		{1, 1},
		{2, 2},
		{3, 3},
	}

	patternColor := color.RGBA{R: 255, G: 255, B: 255, A: 255}

	newImg, err := PaintPatternsOnBlankImage(img, []Pattern{pattern}, patternColor, bgColor)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedColor := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	for _, c := range pattern {
		if newImg.At(c.X, c.Y) != expectedColor {
			t.Errorf("Expected color %v at pixel (%d, %d), but got %v", expectedColor, c.X, c.Y, newImg.At(c.X, c.Y))
		}
	}

	outOfBoundsPattern := Pattern{
		{6, 6},
	}
	_, err = PaintPatternsOnBlankImage(img, []Pattern{outOfBoundsPattern}, patternColor, bgColor)
	expectedErrorMsg := "pixel coordinates out of bounds {6, 6}"
	if err == nil {
		t.Error("Expected an error, but got nil")
	} else if err.Error() != expectedErrorMsg {
		t.Errorf("Expected error message '%s', but got '%s'", expectedErrorMsg, err.Error())
	}
}
