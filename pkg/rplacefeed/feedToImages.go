package rplacefeed

import (
	"encoding/csv"
	"errors"
	"fmt"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func FeedToImages(paths []string) {
	// img := image.NewRGBA(image.Rect(0, 0, 0, 0))

	for _, path := range paths {
		file, err := os.Open(path)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		reader := csv.NewReader(file)

		for {
			record, err := reader.Read()

			if err != nil {
				break
			}

			if len(record) < 3 {
				fmt.Println("Unexpected record with less than 3 columns.")
			} else if len(record) > 4 {
				fmt.Println("Unexpected record with less than 3 columns.")
			}

			// rest here
		}
	}
}

// setPixel sets the color of a pixel at the specified coordinates in the image.
// It checks whether the pixel coordinates are within bounds and returns an error
// including invalid coordinates if they are out of bounds.
func setPixel(img draw.Image, x, y int, color color.Color) error {
	if x < 0 || x >= img.Bounds().Dx() || y < 0 || y >= img.Bounds().Dy() {
		errorMsg := fmt.Sprintf("pixel coordinates out of bounds {%d, %d}", x, y)
		return errors.New(errorMsg)
	}
	img.Set(x, y, color)
	return nil
}
