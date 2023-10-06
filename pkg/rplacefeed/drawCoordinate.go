package rplacefeed

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"math"
)

// DrawCoordinates draws shapes or points on the given image based on the provided coordinates and color.
// It also expands the image and adjusts offsets if necessary to ensure the coordinates fit within the image bounds.
// Returns the modified image, updated offsets, and an error, if any.
func DrawCoordinates(img draw.Image, coordinates Coordinates, color color.Color, offsetLeft int, offsetTop int, verbose bool) (draw.Image, int, int, error) {
	leftExpand, rightExpand, topExpand, bottomExpand := CalculateExpansion(img, coordinates, offsetLeft, offsetTop)

	if leftExpand != 0 || rightExpand != 0 || topExpand != 0 || bottomExpand != 0 {
		img = expandImage(img, topExpand, bottomExpand, leftExpand, rightExpand)
		offsetLeft += leftExpand
		offsetTop += topExpand
	}

	if coordinates.IsPoint() {
		x := coordinates.X + offsetLeft
		y := coordinates.Y + offsetTop

		if verbose {
			fmt.Printf("Point at: %d, %d; color: %+v\n", x, y, color)
		}

		modifiedImg, err := setPixel(img, x, y, color)

		img = modifiedImg
		if err != nil {
			return img, offsetLeft, offsetTop, err
		}
	} else if coordinates.IsCircle() {
		x := coordinates.Circle.X + offsetLeft
		y := coordinates.Circle.Y + offsetTop

		if verbose {
			fmt.Printf("Circle at: %d, %d; radius: %d; color: %+v\n", x, y, coordinates.Circle.R, color)
		}

		modifiedImg := drawCircle(img, x, y, coordinates.Circle.R, color)
		img = modifiedImg
	} else if coordinates.IsRectangle() {
		x1 := coordinates.Rectangle.X1 + offsetLeft
		y1 := coordinates.Rectangle.Y1 + offsetTop
		x2 := coordinates.Rectangle.X2 + offsetLeft
		y2 := coordinates.Rectangle.Y2 + offsetTop

		if verbose {
			fmt.Printf("Rectangle at(x1, y1, x2, y2): %d, %d, %d, %d; color: %+v\n", x1, y1, x2, y2, color)

		}

		modifiedImg := drawRectangle(img, x1, y1, x2, y2, color)
		img = modifiedImg
	} else {
		return img, offsetLeft, offsetTop, errors.New("coordinates don't match point, circle or rectangle ")
	}

	return img, offsetLeft, offsetTop, nil
}

// CalculateExpansion calculates how much the image needs to be expanded on each side (left, right, top, bottom)
// to accommodate the specified coordinates. It takes into account the current offsets.
func CalculateExpansion(img image.Image, coordinate Coordinates, offsetLeft int, offsetTop int) (leftExpand, rightExpand, topExpand, bottomExpand int) {
	imgWidth := img.Bounds().Dx()
	imgHeight := img.Bounds().Dy()

	leftExpand = 0
	rightExpand = 0
	topExpand = 0
	bottomExpand = 0

	if coordinate.IsCircle() {
		circle := coordinate.Circle
		x := circle.X + offsetLeft
		y := circle.Y + offsetTop
		radius := circle.R

		if x-radius < 0 {
			leftExpand = -x + radius
		}
		if x+radius >= imgWidth {
			rightExpand = x + radius - imgWidth + 1
		}
		if y-radius < 0 {
			topExpand = -y + radius
		}
		if y+radius >= imgHeight {
			bottomExpand = y + radius - imgHeight + 1
		}
	} else if coordinate.IsRectangle() {
		rectangle := coordinate.Rectangle
		x1 := rectangle.X1 + offsetLeft
		y1 := rectangle.Y1 + offsetTop
		x2 := rectangle.X2 + offsetLeft
		y2 := rectangle.Y2 + offsetTop

		if x1 < 0 {
			leftExpand = -x1
		}
		if x2 >= imgWidth {
			rightExpand = x2 - imgWidth + 1
		}
		if y1 < 0 {
			topExpand = -y1
		}
		if y2 >= imgHeight {
			bottomExpand = y2 - imgHeight + 1
		}
	} else {
		x := coordinate.X + offsetLeft
		y := coordinate.Y + offsetTop

		if x < 0 {
			leftExpand = -x
		}
		if x >= imgWidth {
			rightExpand = x - imgWidth + 1
		}
		if y < 0 {
			topExpand = -y
		}
		if y >= imgHeight {
			bottomExpand = y - imgHeight + 1
		}
	}

	return leftExpand, rightExpand, topExpand, bottomExpand
}

// expandImage expands the given image by adding empty space around it on all sides, as specified by top, bottom, left, and right values.
// Returns the expanded image.
func expandImage(img draw.Image, top, bottom, left, right int) draw.Image {
	width := img.Bounds().Dx() + left + right
	height := img.Bounds().Dy() + top + bottom

	newImage := image.NewRGBA(image.Rect(0, 0, width, height))

	emptyColor := color.White
	draw.Draw(newImage, newImage.Bounds(), &image.Uniform{emptyColor}, image.Point{}, draw.Over)

	offset := image.Point{left, top}
	draw.Draw(newImage, newImage.Bounds().Add(offset), img, img.Bounds().Min, draw.Over)

	return newImage
}

// setPixel sets the color of a pixel at the specified coordinates (x, y) in the given image.
// Returns the modified image and an error if the coordinates are out of bounds.
func setPixel(img draw.Image, x, y int, color color.Color) (draw.Image, error) {
	if x < 0 || x >= img.Bounds().Dx() || y < 0 || y >= img.Bounds().Dy() {
		errorMsg := fmt.Sprintf("pixel coordinates out of bounds {%d, %d}", x, y)
		return img, errors.New(errorMsg)
	}
	img.Set(x, y, color)

	return img, nil
}

// drawRectangle draws a filled rectangle on the provided image.
// It starts from the top-left corner (x1, y1) and ends at the bottom-right corner (x2, y2),
// filling pixels inside the rectangle with the specified color.
func drawRectangle(img draw.Image, x1, y1, x2, y2 int, color color.Color) draw.Image {
	rect := image.Rect(x1, y1, x2, y2)
	draw.Draw(img, rect, &image.Uniform{color}, image.Point{}, draw.Src)

	return img
}

// drawCircle draws a filled circle on the provided image.
// It starts from the center (x, y) with the given radius,
// filling pixels inside the circle with the specified color.
func drawCircle(img draw.Image, x, y, radius int, color color.Color) draw.Image {
	for i := x - radius; i <= x+radius; i++ {
		for j := y - radius; j <= y+radius; j++ {
			dx := float64(i-x) + 0.5
			dy := float64(j-y) + 0.5

			distance := math.Sqrt(dx*dx + dy*dy)

			if distance <= float64(radius) {
				img.Set(i, j, color)
			}
		}
	}

	return img
}
