package main

import (
	"flag"
	"fmt"
	"image/color"
	"image/png"
	"os"
	"pixel-art-tools/internal/imageloader"
	"pixel-art-tools/pkg/patternmatcher"
	"pixel-art-tools/pkg/rplacefeed"
)

func main() {
	var mode string
	var sourceImagePath string
	var targetImagePath string
	var tolerance int
	var outputFileName string
	var targetColor string

	flag.StringVar(&mode, "mode", "countInstances", "Action that will be performed on provided images")

	flag.StringVar(&sourceImagePath, "sourceImagePath", "", "Path to the source image file")
	flag.StringVar(&targetImagePath, "targetImagePath", "", "Path to the target image file")
	flag.IntVar(&tolerance, "tolerance", 1, "Tolerance for extracting pattern from target image")
	flag.StringVar(&outputFileName, "outputFileName", "visualization.png", "Name of the file generated with visualize function")
	flag.StringVar(&targetColor, "targetColor", "#000000", "Color used for searching for target pattern")

	var directoryPath string
	var baseName string
	var numbersInName int
	var amountOfFiles int
	var verbose bool
	var saveEveryHours bool
	var saveEveryMinutes bool
	var saveEveryValue int
	var outputDir string

	flag.StringVar(&directoryPath, "directoryPath", "./", "Path to diretory containin rplace csv feed")
	flag.StringVar(&baseName, "baseName", "2023_place_canvas_history-", "Full name of the files without numbers")
	flag.IntVar(&numbersInName, "numbersInPath", 12, "Amount of numbers that are present after base name")
	flag.IntVar(&amountOfFiles, "amountOfFiles", 1, "Amount of files that should be used for creating images")
	flag.BoolVar(&verbose, "verbose", false, "Specifies if output should be verbose, verbose output contains extensive logging")
	flag.BoolVar(&saveEveryHours, "saveEveryHours", false, "Specifies if images should be saved in hours intervals")
	flag.BoolVar(&saveEveryMinutes, "saveEveryMinutes", false, "Specifies if images should be saved in minutes intervals")
	flag.IntVar(&saveEveryValue, "saveEveryValue", 1, "Interval of hours or minutes for every which image will be saved")
	flag.StringVar(&outputDir, "outputDit", "output", "Name or full path of directory where pictures will be created. If it doesn't exist then it will be created")

	flag.Parse()

	switch mode {
	case "visualize":
		visualize(sourceImagePath, targetImagePath, tolerance, outputFileName, targetColor)
	case "countInstances":
		countInstances(sourceImagePath, targetImagePath, tolerance, targetColor)
	case "imagesFromRplaceFeed":
		imagesFromRplaceFeed(directoryPath, baseName, numbersInName, amountOfFiles, verbose, saveEveryHours, saveEveryMinutes, saveEveryValue, outputDir)
	}
}

func visualize(sourceImagePath string, targetImagePath string, tolerance int, outputFileName string, targetColor string) {
	if sourceImagePath == "" {
		panic("Path to source image can't be empty")
	}
	if targetImagePath == "" {
		panic("Path to target image can't be empty")
	}
	if tolerance < 0 {
		panic("Tolerance has to be bigger or equal to 0")
	}
	if outputFileName == "" {
		panic("Output file name can't be empty")
	}

	parsedTargetColor, err := rplacefeed.ParseHexColor(targetColor)

	if err != nil {
		panic(fmt.Sprintf("Invalid target color: %s. Please provide color in form of hex code\n", targetColor))
	}

	mainImage, err := imageloader.LoadImage(sourceImagePath)
	if err != nil {
		panic(err)
	}

	targetImage, err := imageloader.LoadImage(targetImagePath)
	if err != nil {
		panic(err)
	}

	pattern, err := patternmatcher.ImageToPattern(targetImage, parsedTargetColor, 5)

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

func countInstances(sourceImagePath string, targetImagePath string, tolerance int, targetColor string) {
	if sourceImagePath == "" {
		panic("Path to source image can't be empty")
	}
	if targetImagePath == "" {
		panic("Path to target image can't be empty")
	}
	if tolerance < 0 {
		panic("Tolerance has to be bigger or equal to 0")
	}

	parsedTargetColor, err := rplacefeed.ParseHexColor(targetColor)

	if err != nil {
		panic(fmt.Sprintf("Invalid target color: %s. Please provide color in form of hex code\n", targetColor))
	}

	sourceImage, err := imageloader.LoadImage(sourceImagePath)
	if err != nil {
		panic(err)
	}

	targetImage, err := imageloader.LoadImage(targetImagePath)
	if err != nil {
		panic(err)
	}

	pattern, err := patternmatcher.ImageToPattern(targetImage, parsedTargetColor, 5)

	if err != nil {
		panic(err)
	}

	patternsInImage := patternmatcher.PatternsInImage(sourceImage, pattern)

	fmt.Println(len(patternsInImage))
}

func imagesFromRplaceFeed(directoryPath string, baseName string, numbersInName int, amountOfFiles int, verbose bool, saveEveryHours bool, saveEveryMinutes bool, saveEveryValue int, outputDir string) {
	if directoryPath == "" {
		panic("Directory path can't be empty")
	}
	if numbersInName <= 0 {
		panic("Numbers in name has to be bigger than 0")
	}
	if amountOfFiles <= 0 {
		panic("Amount of files has to be bigger than 0")
	}
	if !saveEveryHours && !saveEveryMinutes {
		panic("Program in this mode requires saveEveryHours or saveEveryMinutes set to true")
	}
	if saveEveryHours && saveEveryMinutes {
		panic("Both saveEveryHours and saveEveryMinutes were set to true. Please choose only one")
	}
	if saveEveryValue <= 0 {
		panic("saveEveryValue has to be bigger than 0")
	}
	if outputDir == "" {
		panic("Output directory can't be empty")
	}

	fileNames := rplacefeed.GenerateFileNames(baseName, numbersInName, amountOfFiles, ".csv")
	paths := rplacefeed.GenerateAndVerifyPaths(directoryPath, fileNames)

	options := rplacefeed.FeedToImagesOptions{Verbose: verbose, SaveEveryHours: saveEveryHours, SaveEveryMinutes: saveEveryMinutes, SaveEveryValue: saveEveryValue, OutputDir: outputDir}
	_, err := rplacefeed.FeedToImages(paths, options)

	if err != nil {
		panic(err)
	}
}
