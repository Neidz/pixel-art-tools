package main

import (
	"flag"
	"fmt"
	"pixel-art-tools/pkg/rplacefeed"
)

func main() {
	baseName := "2023_place_canvas_history-"
	var directoryPath string

	flag.StringVar(&directoryPath, "directoryPath", "./", "Path to diretory containin rplace csv feed")

	flag.Parse()

	numbersInPath := 12
	// amountOfFiles := 53
	amountOfFiles := 5

	fileNames := rplacefeed.GenerateFileNames(baseName, numbersInPath, amountOfFiles, ".csv")

	paths := rplacefeed.GenerateAndVerifyPaths(directoryPath, fileNames)

	fmt.Println(paths)
}
