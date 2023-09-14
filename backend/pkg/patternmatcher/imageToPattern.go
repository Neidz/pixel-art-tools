package patternmatcher

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
)

func ImageToPattern(img image.Image, searchedColor color.Color, tolerance int) (Pattern, error) {
	imgWidth := img.Bounds().Dx()
	imgHeight := img.Bounds().Dy()

	var pattern Pattern

	for y := 0; y < imgHeight; y++ {
		for x := 0; x < imgWidth; x++ {
			isSearchedColor, err := AreColorsSimilar(img.At(x, y), searchedColor, tolerance)

			if err != nil {
				return pattern, err
			}
			if isSearchedColor {
				pattern = append(pattern, Coordinate{x, y})
			}
		}
	}

	return pattern, nil
}
