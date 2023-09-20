package rplacefeed

import "strconv"

func GenerateFileNames(baseName string, numbersInName int, amountOfFiles int, extension string) []string {
	var fileNames []string

	for i := 0; i < amountOfFiles; i++ {
		fileNames = append(fileNames, generateFileName(baseName, numbersInName, i, extension))
	}

	return fileNames
}

func generateFileName(baseName string, numbersInName int, number int, extension string) string {
	intStr := strconv.Itoa(number)

	fileName := baseName

	zeros := numbersInName - len(intStr)

	for i := 0; i < zeros; i++ {
		fileName += "0"
	}

	fileName += intStr
	fileName += extension

	return fileName
}
