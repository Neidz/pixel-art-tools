package patternmatcher

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
)

// PaintPatternsOnBlankImage paints patterns on a blank image and returns the resulting image.
// It takes an original image, a list of patterns to paint, the pattern color, and the background color.
// It returns the resulting image or an error if pixel coordinates are out of bounds.
func PaintPatternsOnBlankImage(img image.Image, patterns []Pattern, patternColor color.Color, bgColor color.Color) (image.Image, error) {
	imgWidth := img.Bounds().Dx()
	imgHeight := img.Bounds().Dy()

	newImg := filledImage(imgWidth, imgHeight, bgColor)

	for _, p := range patterns {
		for _, c := range p {
			err := setPixel(newImg, c.X, c.Y, patternColor)
			if err != nil {
				return nil, err
			}
		}
	}

	return newImg, nil
}

// filledImage creates a new filled image with the specified dimensions and background color.
// It returns the newly created image.
func filledImage(width int, height int, bgColor color.Color) draw.Image {
	rect := image.Rectangle{image.Point{0, 0}, image.Point{width, height}}

	img := image.NewRGBA(rect)

	draw.Draw(img, rect, &image.Uniform{bgColor}, image.Point{}, draw.Src)

	return img
}

// setPixel sets the color of a pixel at the specified coordinates in the image.
// It checks whether the pixel coordinates are within bounds and returns an error
// including invalig coordinates if they are out of bounds.
func setPixel(img draw.Image, x, y int, color color.Color) error {
	if x < 0 || x >= img.Bounds().Dx() || y < 0 || y >= img.Bounds().Dy() {
		errorMsg := fmt.Sprintf("pixel coordinates out of bounds {%d, %d}", x, y)
		return errors.New(errorMsg)
	}
	img.Set(x, y, color)
	return nil
}
