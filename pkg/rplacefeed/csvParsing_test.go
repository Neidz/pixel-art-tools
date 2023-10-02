package rplacefeed

import (
	"testing"
)

func TestParseRecord(t *testing.T) {
	testCases := []struct {
		name                string
		coordinatesStr      string
		expectedCoordinates Coordinates
	}{
		{
			name:                "Valid Simple Coordinate",
			coordinatesStr:      "\"364,295\"",
			expectedCoordinates: Coordinates{X: 364, Y: 295},
		},
		{
			name:                "Valid Circle Coordinate",
			coordinatesStr:      "\"{X:424,Y:336,R:3}\"",
			expectedCoordinates: Coordinates{X: 0, Y: 0, Circle: Circle{X: 424, Y: 336, R: 3}},
		},
		{
			name:                "Valid Rectangle Coordinate",
			coordinatesStr:      "\"20,30,40,50\"",
			expectedCoordinates: Coordinates{X: 0, Y: 0, Rectangle: Rectangle{X1: 20, Y1: 30, X2: 40, Y2: 50}},
		},
		{
			name:                "Valid Negative Coordinate",
			coordinatesStr:      "\"364,-295\"",
			expectedCoordinates: Coordinates{X: 364, Y: -295},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCoordinates(t, testCase.coordinatesStr, testCase.expectedCoordinates)
		})
	}
}

func testCoordinates(t *testing.T, coordinatesStr string, expectedCoordinates Coordinates) {
	date := "2023-07-20 13:06:39.31 UTC"
	user := "8OV068iJZlClecOX+RO3xeOhkim9VgFRZAw9IUfqmYCrJShk9d/CPl/D6NOTXEZzk6n5oTuQSc8revdnc1GXvg=="
	color := "#FF4500"
	record, err := ParseRecord(date, user, coordinatesStr, color)

	if err != nil {
		t.Error("Error: ", err)
	}

	if record.Coordinates != expectedCoordinates {
		t.Errorf("Expected coordinates %+v but got %+v", expectedCoordinates, record.Coordinates)
	}
}
