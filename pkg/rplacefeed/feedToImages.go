package rplacefeed

import (
	"encoding/csv"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
)

func FeedToImages(paths []string) (image.Image, error) {
	img := image.NewRGBA(image.Rect(0, 0, 1, 1))
	offsetLeft := 0
	offsetTop := 0

	for i, path := range paths {
		recordNumber := 0
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		reader := csv.NewReader(file)

		for {
			record, err := reader.Read()
			recordNumber++

			fmt.Println("Record: ", i, ", line: ", recordNumber)

			if err == io.EOF {
				fmt.Println("Finished parsing file with path: ", path)
				break
			} else if err != nil {
				return nil, err
			}

			if len(record) < 3 {
				fmt.Println("Unexpected record with less than 3 columns.")
				continue
			} else if len(record) > 4 {
				fmt.Println("Unexpected record with less than 3 columns.")
				continue
			}

			parsedRecord, err := ParseRecord(record[0], record[1], record[2], record[3])

			if err != nil {
				fmt.Println("Error: ", err)
				continue
			}

			modifiedImg, modifiedOffsetLeft, modifiedOffsetTop, drawErr := DrawCoordinates(img, parsedRecord.Coordinates, parsedRecord.Color, offsetLeft, offsetTop)

			if drawErr != nil {
				fmt.Println("Error: ", drawErr)
			} else {
				modifiedImg, ok := modifiedImg.(*image.RGBA)

				if !ok {
					fmt.Println("Error: failed to convert image.Image to *image.RGBA")
				} else {
					img = modifiedImg
					offsetLeft = modifiedOffsetLeft
					offsetTop = modifiedOffsetTop
				}
			}
		}
	}

	return img, nil
}
