package rplacefeed

import (
	"image/color"
	"testing"
)

func TestParseRecord(t *testing.T) {
	testCases := []struct {
		name                string
		colorStr            string
		expectedColor       color.Color
		coordinatesStr      string
		expectedCoordinates Coordinates
	}{
		{
			name:                "Valid Simple Coordinate",
			colorStr:            "#FFFFFF",
			expectedColor:       color.RGBA{R: 255, G: 255, B: 255, A: 255},
			coordinatesStr:      "\"364,295\"",
			expectedCoordinates: Coordinates{X: 364, Y: 295},
		},
		{
			name:                "Valid Circle Coordinate",
			colorStr:            "#FF0000",
			expectedColor:       color.RGBA{R: 255, G: 0, B: 0, A: 255},
			coordinatesStr:      "\"{X:424,Y:336,R:3}\"",
			expectedCoordinates: Coordinates{X: 0, Y: 0, Circle: Circle{X: 424, Y: 336, R: 3}},
		},
		{
			name:                "Valid Rectangle Coordinate",
			colorStr:            "#00FF00",
			expectedColor:       color.RGBA{R: 0, G: 255, B: 0, A: 255},
			coordinatesStr:      "\"20,30,40,50\"",
			expectedCoordinates: Coordinates{X: 0, Y: 0, Rectangle: Rectangle{X1: 20, Y1: 30, X2: 40, Y2: 50}},
		},
		{
			name:                "Valid Negative Coordinate",
			colorStr:            "#0000FF",
			expectedColor:       color.RGBA{R: 0, G: 0, B: 255, A: 255},
			coordinatesStr:      "\"364,-295\"",
			expectedCoordinates: Coordinates{X: 364, Y: -295},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCoordinates(t, testCase.colorStr, testCase.expectedColor, testCase.coordinatesStr, testCase.expectedCoordinates)
		})
	}
}

func testCoordinates(t *testing.T, colorStr string, expectedColor color.Color, coordinatesStr string, expectedCoordinates Coordinates) {
	date := "2023-07-20 13:06:39.31 UTC"
	user := "8OV068iJZlClecOX+RO3xeOhkim9VgFRZAw9IUfqmYCrJShk9d/CPl/D6NOTXEZzk6n5oTuQSc8revdnc1GXvg=="
	record, err := ParseRecord(date, user, coordinatesStr, colorStr)

	if err != nil {
		t.Error("Error: ", err)
	}

	if record.Color != expectedColor {
		t.Errorf("Expected color %+v but got %+v", expectedColor, record.Color)
	}

	if record.Coordinates != expectedCoordinates {
		t.Errorf("Expected coordinates %+v but got %+v", expectedCoordinates, record.Coordinates)
	}
}
