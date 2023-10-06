package main

import (
	"flag"
	"pixel-art-tools/pkg/rplacefeed"
)

func main() {
	baseName := "2023_place_canvas_history-"
	var directoryPath string

	flag.StringVar(&directoryPath, "directoryPath", "./", "Path to diretory containin rplace csv feed")

	flag.Parse()

	numbersInPath := 12
	amountOfFiles := 5

	fileNames := rplacefeed.GenerateFileNames(baseName, numbersInPath, amountOfFiles, ".csv")
	paths := rplacefeed.GenerateAndVerifyPaths(directoryPath, fileNames)

	options := rplacefeed.FeedToImagesOptions{Verbose: false, SaveEveryHours: true, SaveEveryMinutes: false, SaveEveryValue: 1, OutputDir: "output"}
	_, err := rplacefeed.FeedToImages(paths, options)

	if err != nil {
		panic(err)
	}
}
