package patternmatcher

import (
	"image/color"
	"testing"
)

func TestAreColorsSimilar(t *testing.T) {
	color1 := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	color2 := color.RGBA{R: 250, G: 255, B: 255, A: 255}
	color3 := color.RGBA{R: 249, G: 255, B: 255, A: 255}

	result1, err1 := AreColorsSimilar(color1, color2, 5)
	if err1 != nil {
		t.Errorf("Unexpected error: %v", err1)
	}

	if !result1 {
		t.Errorf("AreColorsSimilar should return true for %v and %v with tolerance 5", color1, color2)
	}

	result2, err2 := AreColorsSimilar(color1, color3, 5)
	if err2 != nil {
		t.Errorf("Unexpected error: %v", err2)
	}

	if result2 {
		t.Errorf("AreColorsSimilar should return false for %v and %v with tolerance 5", color1, color3)
	}
}

func TestAbsInt(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{5, 5},
		{-5, 5},
		{0, 0},
	}

	for _, tc := range testCases {
		result := absInt(tc.input)
		if result != tc.expected {
			t.Errorf("absInt(%d) = %d; want %d", tc.input, result, tc.expected)
		}
	}
}
