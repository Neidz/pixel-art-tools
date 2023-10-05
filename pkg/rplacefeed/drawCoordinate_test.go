package rplacefeed

import (
	"image"
	"image/color"
	"image/draw"
	"math"
	"testing"
)

func TestSetPixel(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))

	x := 50
	y := 50
	pixelColor := color.RGBA{255, 255, 255, 255}

	resultImg, err := setPixel(img, x, y, pixelColor)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if resultImg.At(x, y) != pixelColor {
		t.Errorf("Pixel color at (%d, %d) is %v, expected %v", x, y, resultImg.At(x, y), pixelColor)
	}

	_, err = setPixel(img, -1, -1, pixelColor)
	if err == nil {
		t.Errorf("Expected an error for out-of-bounds coordinates, but got none")
	}
}

func TestDrawRectangle(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{255, 255, 255, 255}}, image.Point{}, draw.Src)

	x1, y1 := 20, 20
	x2, y2 := 80, 80
	rectangleColor := color.RGBA{255, 255, 255, 255}

	resultImg := drawRectangle(img, x1, y1, x2, y2, rectangleColor)

	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			pixelColor := resultImg.At(i, j)

			if pixelColor != rectangleColor {
				t.Errorf("Pixel color at (%d, %d) is %v, expected %v", i, j, pixelColor, rectangleColor)
			}
		}
	}
}

func TestDrawCircle(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))

	x := 50
	y := 50
	radius := 20
	circleColor := color.RGBA{255, 255, 255, 255}

	resultImg := drawCircle(img, x, y, radius, circleColor)

	for i := x - radius; i <= x+radius; i++ {
		for j := y - radius; j <= y+radius; j++ {
			dx := float64(i-x) + 0.5
			dy := float64(j-y) + 0.5

			distance := math.Sqrt(dx*dx + dy*dy)

			expectedInside := distance <= float64(radius)

			pixelColor := resultImg.At(i, j)

			if expectedInside && pixelColor != circleColor {
				t.Errorf("Pixel color at (%d, %d) is %v, expected %v", i, j, pixelColor, circleColor)
			}
		}
	}
}
