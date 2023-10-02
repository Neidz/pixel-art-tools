package rplacefeed

import "testing"

func TestGenerateFileName(t *testing.T) {
	baseName := "aperture-science-"
	numbersInName := 10
	extension := ".csv"
	number1 := 7
	number2 := 137

	name1 := generateFileName(baseName, numbersInName, number1, extension)
	expectedName1 := "aperture-science-0000000007.csv"

	if name1 != expectedName1 {
		t.Errorf("Expected file name %q but got %q", expectedName1, name1)
	}

	name2 := generateFileName(baseName, numbersInName, number2, extension)
	expectedName2 := "aperture-science-0000000137.csv"

	if name2 != expectedName2 {
		t.Errorf("Expected file name %q but got %q", expectedName2, name2)
	}
}
