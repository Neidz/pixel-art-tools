package main

import (
	"flag"
	"fmt"
	"image/color"
	"image/png"
	"os"
	"pixel-art-tools/backend/pkg/imageloader"
	"pixel-art-tools/backend/pkg/patternmatcher"
)

func main() {
	var action string
	var sourceImagePath string
	var targetImagePath string
	var tolerance int
	var outputFileName string

	flag.StringVar(&action, "action", "countInstances", "Action that will be performed on provided images")
	flag.StringVar(&sourceImagePath, "sourceImagePath", "", "Path to the source image file")
	flag.StringVar(&targetImagePath, "targetImagePath", "", "Path to the target image file")
	flag.IntVar(&tolerance, "tolerance", 1, "Tolerance for extracting pattern from target image")
	flag.StringVar(&outputFileName, "outputFileName", "visualization.png", "Name of the file generated with visualize function")

	flag.Parse()

	if sourceImagePath == "" {
		panic("Provide path to source image")
	}
	if targetImagePath == "" {
		panic("Provide path to target image")
	}

	switch action {
	case "visualize":
		visualize(sourceImagePath, targetImagePath, tolerance, outputFileName)
	case "countInstances":
		countInstances(sourceImagePath, targetImagePath, tolerance)
	}

}

func visualize(sourceImagePath string, targetImagePath string, tolerance int, outputFileName string) {
	mainImage, err := imageloader.LoadImage(sourceImagePath)
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

	file, err := os.Create(outputFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, visualization)
	if err != nil {
		panic(err)
	}
}

func countInstances(sourceImagePath string, targetImagePath string, tolerance int) {
	mainImage, err := imageloader.LoadImage(sourceImagePath)
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

	fmt.Println(len(patternsInImage))
}
