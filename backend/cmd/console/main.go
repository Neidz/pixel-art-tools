package main

import (
	"image/color"
	"image/png"
	"os"
	"pixel-art-tools/backend/pkg/imageloader"
	"pixel-art-tools/backend/pkg/patternmatcher"
)

func main() {
	mainImagePath := "/home/neidz/Projects/pixel-art-tools/backend/assets/final_2023_place.png"
	targetImagePath := "/home/neidz/Projects/pixel-art-tools/backend/assets/reversedCrewmate.png"

	mainImage, err := imageloader.LoadImage(mainImagePath)
	if err != nil {
		panic(err)
	}

	targetImage, err := imageloader.LoadImage(targetImagePath)
	if err != nil {
		panic(err)
	}

	pattern, err := patternmatcher.ImageToPattern(targetImage, color.White, 5)

	if err != nil {
		panic(err)
	}

	patternsInImage := patternmatcher.PatternsInImage(mainImage, pattern)
	visualization, err := patternmatcher.PaintPatternsOnBlankImage(mainImage, patternsInImage, color.White, color.Black)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("visualization.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Encode and save the image as a PNG
	err = png.Encode(file, visualization)
	if err != nil {
		panic(err)
	}
}
