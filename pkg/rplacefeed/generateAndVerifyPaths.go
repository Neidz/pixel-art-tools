package rplacefeed

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func GenerateAndVerifyPaths(basePath string, fileNames []string) []string {
	var paths []string

	for _, fileName := range fileNames {
		path := filepath.Join(basePath, fileName)
		lineCount, err := verifyFile(path)
		fmt.Println("Path: ", path)

		if err != nil {
			fmt.Println("Error: ", err)
		} else {
			paths = append(paths, path)
			fmt.Printf("Found file with name %v containing %v lines.", fileName, lineCount)
		}
	}

	return paths
}

func verifyFile(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return -1, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineCount := 0

	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return -1, err
	}

	return lineCount, nil
}
