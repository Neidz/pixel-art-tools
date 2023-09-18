package patternmatcher

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"sync"
)

// PatternsInImage searches for a specified pattern within an image and returns a slice
// containing all occurrences of the pattern as patterns. Pattern is recognized as valid
// only if all of it's coordinates contain exactly the same color and if every coordinate
// around pattern coordinates that are not part coordinates have different color.
func PatternsInImage(img image.Image, searchedPattern Pattern) []Pattern {
	tolerance := 0
	imgWidth := img.Bounds().Dx()
	imgHeight := img.Bounds().Dy()

	var patterns []Pattern
	var wg sync.WaitGroup

	results := make(chan Pattern)

	spWidth, spHeight := patternBoxSize(searchedPattern)

	// Iterate through the image to search for the pattern.
	for offsetY := 0; offsetY <= imgHeight-spHeight; offsetY++ {
		for offsetX := 0; offsetX <= imgWidth-spWidth; offsetX++ {
			wg.Add(1)
			go func(offsetX int, offsetY int) {
				defer wg.Done()
				pattern := processWindow(img, searchedPattern, offsetX, offsetY, tolerance)
				if pattern != nil {
					results <- pattern
				}
			}(offsetX, offsetY)
		}
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for pattern := range results {
		patterns = append(patterns, pattern)
	}

	return patterns
}

// processWindow checks if a pattern exists within a specific window of the image
// defined by the provided offset values. If the pattern is found within the window
// it returns the pattern with adjusted coordinates relative to the window.
// If the pattern is not found, it returns nil.
func processWindow(img image.Image, sp Pattern, offsetX, offsetY, tolerance int) Pattern {
	if isPatternInWindow(img, sp, offsetX, offsetY, tolerance) {
		var pattern Pattern

		for _, coordinate := range sp {
			pattern = append(pattern, Coordinate{coordinate.X + offsetX, coordinate.Y + offsetY})
		}

		return pattern
	}
	return nil
}

// patternBoxSize calculates the width and height of smallest box that would contain provided pattern.
func patternBoxSize(p Pattern) (int, int) {
	highestX := 0
	highestY := 0

	for _, coordinate := range p {
		if coordinate.X > highestX {
			highestX = coordinate.X
		}
		if coordinate.Y > highestY {
			highestY = coordinate.Y
		}
	}

	return highestX + 1, highestY + 1
}

// isPatternInWindow checks if a pattern exists in a specific window of the image.
func isPatternInWindow(img image.Image, sp Pattern, offsetX int, offsetY int, tolerance int) bool {
	firstPixelColor := img.At(sp[0].X+offsetX, sp[1].Y+offsetY)

	for _, coordinate := range sp {
		coordinateWithOffset := Coordinate{coordinate.X + offsetX, coordinate.Y + offsetY}

		colorMatch, err := AreColorsSimilar(firstPixelColor, img.At(coordinateWithOffset.X, coordinateWithOffset.Y), tolerance)
		if err != nil {
			fmt.Println("Error: ", err)
			return false
		}
		if !colorMatch {
			return false
		}

		for _, surroundingOffset := range surroundingOffsets {
			pixelLocation := Coordinate{coordinateWithOffset.X + surroundingOffset.X, coordinateWithOffset.Y + surroundingOffset.Y}

			if pixelLocation.X < 0 || pixelLocation.Y < 0 || pixelLocation.X >= img.Bounds().Dx() || pixelLocation.Y >= img.Bounds().Dy() {
				continue
			}

			inPattern := containsElement(sp, Coordinate{pixelLocation.X - offsetX, pixelLocation.Y - offsetY})
			sameColor, err := AreColorsSimilar(firstPixelColor, img.At(pixelLocation.X, pixelLocation.Y), tolerance)
			if err != nil {
				fmt.Println("Error: ", err)
				return false
			}

			if !inPattern && sameColor {
				return false
			}
		}
	}

	return true
}

func containsElement(slice Pattern, el Coordinate) bool {
	for _, item := range slice {
		if item == el {
			return true
		}
	}
	return false
}

const (
	left   = -1
	right  = 1
	top    = -1
	bottom = 1
	center = 0
)

// surroundingOffsets represents a slice of Coordinate offsets for surrounding pixels.
var surroundingOffsets = []Coordinate{{left, top}, {center, top}, {right, top}, {left, center}, {right, center}, {left, bottom}, {center, bottom}, {right, bottom}}
