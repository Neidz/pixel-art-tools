package patternmatcher

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
)

// ImageToPattern converts an image into a Pattern by searching for pixels that
// match a specified color within a given tolerance. It returns the Pattern and
// an error if there is an issue with color comparison or image processing.
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
