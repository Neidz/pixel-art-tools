package main

import (
	"flag"
	"image/png"
	"os"
	"pixel-art-tools/pkg/rplacefeed"
)

func main() {
	baseName := "2023_place_canvas_history-"
	var directoryPath string

	flag.StringVar(&directoryPath, "directoryPath", "./", "Path to diretory containin rplace csv feed")

	flag.Parse()

	numbersInPath := 12

	amountOfFiles := 1

	fileNames := rplacefeed.GenerateFileNames(baseName, numbersInPath, amountOfFiles, ".csv")

	paths := rplacefeed.GenerateAndVerifyPaths(directoryPath, fileNames)

	img, err := rplacefeed.FeedToImages(paths)

	if err != nil {
		panic(err)
	}

	file, err := os.Create("rplaceFeedImage")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}
