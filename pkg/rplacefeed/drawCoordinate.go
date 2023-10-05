package rplacefeed

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
)

func DrawCoordinates(img draw.Image, coordinates Coordinates, color color.Color, offsetLeft int, offsetTop int) (draw.Image, int, int, error) {
	leftExpand, rightExpand, topExpand, bottomExpand := CalculateExpansion(img, coordinates, offsetLeft, offsetTop)

	if leftExpand != 0 || rightExpand != 0 || topExpand != 0 || bottomExpand != 0 {
		img = expandImage(img, topExpand, bottomExpand, leftExpand, rightExpand)
		offsetLeft += leftExpand
		offsetTop += topExpand
	}

	if coordinates.IsPoint() {
		x := coordinates.X + offsetLeft
		y := coordinates.Y + offsetTop

		fmt.Printf("Point at: %d, %d; color: %+v\n", x, y, color)

		modifiedImg, err := setPixel(img, x, y, color)

		img = modifiedImg
		if err != nil {
			return img, offsetLeft, offsetTop, err
		}
	} else if coordinates.IsCircle() {
		x := coordinates.Circle.X + offsetLeft
		y := coordinates.Circle.Y + offsetTop

		fmt.Printf("Circle at: %d, %d; radius: %d; color: %+v\n", x, y, coordinates.Circle.R, color)

		modifiedImg := drawCircle(img, x, y, coordinates.Circle.R, color)
		img = modifiedImg
	} else if coordinates.IsRectangle() {
		x1 := coordinates.Rectangle.X1 + offsetLeft
		y1 := coordinates.Rectangle.Y1 + offsetTop
		x2 := coordinates.Rectangle.X2 + offsetLeft
		y2 := coordinates.Rectangle.Y2 + offsetTop

		fmt.Printf("Rectangle at(x1, y1, x2, y2): %d, %d, %d, %d; color: %+v\n", x1, y1, x2, y2, color)

		modifiedImg := drawRectangle(img, x1, y1, x2, y2, color)
		img = modifiedImg
	} else {
		return img, offsetLeft, offsetTop, errors.New("coordinates don't match point, circle or rectangle ")
	}

	return img, offsetLeft, offsetTop, nil
}

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

func setPixel(img draw.Image, x, y int, color color.Color) (draw.Image, error) {
	if x < 0 || x >= img.Bounds().Dx() || y < 0 || y >= img.Bounds().Dy() {
		errorMsg := fmt.Sprintf("pixel coordinates out of bounds {%d, %d}", x, y)
		return img, errors.New(errorMsg)
	}
	img.Set(x, y, color)

	return img, nil
}

func drawRectangle(img draw.Image, x1, y1, x2, y2 int, color color.Color) draw.Image {
	rect := image.Rect(x1, y1, x2, y2)
	draw.Draw(img, rect, &image.Uniform{color}, image.Point{}, draw.Src)

	return img
}

func drawCircle(img draw.Image, x, y, radius int, color color.Color) draw.Image {
	draw.DrawMask(img, img.Bounds(), &image.Uniform{color}, image.Point{}, &circleMask{x, y, radius}, image.Point{}, draw.Over)

	return img
}

type circleMask struct {
	x, y, r int
}

func (c *circleMask) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *circleMask) Bounds() image.Rectangle {
	return image.Rect(c.x-c.r, c.y-c.r, c.x+c.r, c.y+c.r)
}

func (c *circleMask) At(x, y int) color.Color {
	xx, yy, rr := float64(x-c.x)+0.5, float64(y-c.y)+0.5, float64(c.r)
	if xx*xx+yy*yy < rr*rr {
		return color.Alpha{255}
	}
	return color.Alpha{0}
}
