package rplacefeed

import (
	"encoding/csv"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"io"
	"os"
	"path/filepath"
	"time"
)

// FeedToImages reads a list of file paths, parses CSV records from these files, and applies the actions described by the records on an image.
// It returns the final modified image.
//
// Parameters:
// - paths: []string, a list of file paths containing CSV records to process.
// - options: FeedToImagesOptions, configuration options for the processing.
//
// Returns:
// - image.Image: The modified image after applying actions from CSV records.
// - error: An error if any issues occur during processing.
//
// The function initializes an image with a white background and processes CSV records sequentially from the input files.
// It tracks the start date and time, updates action time based on time intervals, and applies pixel modifications to the image.
// The resulting modified image is returned as the output.
func FeedToImages(paths []string, options FeedToImagesOptions) (image.Image, error) {
	img := image.NewRGBA(image.Rect(0, 0, 1, 1))
	offsetLeft := 0
	offsetTop := 0

	if options.SaveEveryMinutes || options.SaveEveryHours {
		createOutputDir(options.OutputDir)
	}

	var startDate time.Time
	dateAssigned := false
	var actionTime time.Time

	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{255, 255, 255, 255}}, image.Point{}, draw.Src)

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

			if options.Verbose {
				fmt.Println("Record: ", i, ", line: ", recordNumber)
			}

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

			if record[0] == "timestamp" {
				continue
			}

			parsedRecord, err := ParseRecord(record[0], record[1], record[2], record[3])

			if err != nil {
				fmt.Println("Error: ", err)
				fmt.Println("Info: timestamp was", record[3])
				continue
			}

			if !dateAssigned {
				startDate = parsedRecord.Date
				actionTime = parsedRecord.Date
				dateAssigned = true
			}

			newActionTime := timeBasedAction(options, img, actionTime, startDate, parsedRecord.Date)

			if newActionTime != actionTime {
				actionTime = newActionTime
			}

			modifiedImg, modifiedOffsetLeft, modifiedOffsetTop, drawErr := DrawCoordinates(img, parsedRecord.Coordinates, parsedRecord.Color, offsetLeft, offsetTop, options.Verbose)

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

// timeBasedAction performs actions based on time intervals and updates the actionTime.
// Parameters:
// - options: FeedToImagesOptions, configuration options.
// - img: *image.RGBA, the image to process.
// - actionTime: time.Time, last action time.
// - firstDate: time.Time, initial recording start time.
// - recordTime: time.Time, current time.
// Returns the updated actionTime.
func timeBasedAction(options FeedToImagesOptions, img *image.RGBA, actionTime time.Time, firstDate time.Time, recordTime time.Time) time.Time {
	timePassed := recordTime.Sub(firstDate)

	hoursSinceStart := int(timePassed.Hours())
	minutesSinceStart := int(timePassed.Minutes()) % 60

	if options.Verbose {
		fmt.Printf("Time passed: %d hours, %d minutes\n", hoursSinceStart, minutesSinceStart)
	}

	if options.SaveEveryHours {
		hoursSinceAction := int(recordTime.Sub(actionTime).Hours())

		if hoursSinceAction >= options.SaveEveryValue {
			saveImage(img, options.OutputDir, formatTime(hoursSinceStart, minutesSinceStart), options.Verbose)
			actionTime = recordTime
		}
	}

	if options.SaveEveryMinutes {
		minutesSinceAction := int(recordTime.Sub(actionTime).Minutes()) % 60

		if minutesSinceAction >= options.SaveEveryValue {
			saveImage(img, options.OutputDir, formatTime(hoursSinceStart, minutesSinceStart), options.Verbose)
			actionTime = recordTime
		}
	}

	return actionTime
}

func formatTime(hours int, minutes int) string {
	return fmt.Sprintf("%d-hours,-%d-minutes", hours, minutes)
}

func saveImage(img image.Image, directory string, information string, verbose bool) {
	name := "rplace-" + information + ".png"
	file, err := os.Create(filepath.Join(directory, name))
	if err != nil {
		fmt.Printf("Failed to create image at '%s' with name '%s'\n", directory, name)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		fmt.Printf("Failed to save image: %s\n", name)
	}

	if verbose {
		fmt.Printf("Image created: %s\n", filepath.Join(directory, name))
	}
}

func createOutputDir(outputDir string) {
	_, err := os.Stat(outputDir)

	if os.IsNotExist(err) {
		err := os.MkdirAll(outputDir, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
			return
		}

		fmt.Printf("Created directory '%s'\n", outputDir)
	} else if err != nil {
		fmt.Printf("Error checking directory: %v\n", err)
		return
	} else {
		fmt.Printf("Directory '%s' already exists\n", outputDir)
	}
}
