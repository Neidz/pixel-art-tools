package rplacefeed

import "testing"

func TestGenerateFileName(t *testing.T) {
	baseName := "aperture-science-"
	numbersInName := 10
	extension := ".csv"
	number1 := 7
	number2 := 137

	path1 := generateFileName(baseName, numbersInName, number1, extension)
	expectedPath1 := "aperture-science-0000000007"

	if path1 != expectedPath1 {
		t.Errorf("Expected file name %q but got %q", expectedPath1, path1)
	}

	path2 := generateFileName(baseName, numbersInName, number2, extension)
	expectedPath2 := "aperture-science-0000000137"

	if path2 != expectedPath2 {
		t.Errorf("Expected file name %q but got %q", expectedPath2, path2)
	}
}
