package rplacefeed

import (
	"errors"
	"fmt"
	"image/color"
	"regexp"
	"strconv"
	"time"
)

// ParseRecord parses a CSV record, extracting date, user, coordinates, and color.
// It returns a CSVRecord struct and an error. If parsing is successful, the error is nil.
// Input Parameters:
//   - dateStr: Timestamp in "2006-01-02 15:04:05.999 UTC" format.
//   - userStr: User identifier.
//   - coordinatesStr: Coordinates in various formats (simple, rectangle, or circle).
//   - colorStr: Hex color associated with the record.
func ParseRecord(dateStr string, userStr, coordinatesStr string, colorStr string) (CSVRecord, error) {
	var record CSVRecord

	timestamp, err := time.Parse("2006-01-02 15:04:05.999 UTC", dateStr)

	if err != nil {
		return CSVRecord{}, err
	}

	color, err := parseHexColor(colorStr)

	if err != nil {
		return record, err
	}

	rectanglePattern := regexp.MustCompile(`(-?\d+),(-?\d+),(-?\d+),(-?\d+)`)
	circlePattern := regexp.MustCompile(`{X:(-?\d+),Y:(-?\d+),R:(-?\d+)}`)
	simplePattern := regexp.MustCompile(`(-?\d+),(-?\d+)`)

	if matches := rectanglePattern.FindStringSubmatch(coordinatesStr); len(matches) == 5 {
		x1, _ := strconv.Atoi(matches[1])
		y1, _ := strconv.Atoi(matches[2])
		x2, _ := strconv.Atoi(matches[3])
		y2, _ := strconv.Atoi(matches[4])
		record.Coordinates.Rectangle = Rectangle{X1: x1, Y1: y1, X2: x2, Y2: y2}
	} else if matches := circlePattern.FindStringSubmatch(coordinatesStr); len(matches) == 4 {
		x, _ := strconv.Atoi(matches[1])
		y, _ := strconv.Atoi(matches[2])
		r, _ := strconv.Atoi(matches[3])
		record.Coordinates.Circle = Circle{X: x, Y: y, R: r}
	} else if matches := simplePattern.FindStringSubmatch(coordinatesStr); len(matches) == 3 {
		x, _ := strconv.Atoi(matches[1])
		y, _ := strconv.Atoi(matches[2])
		record.Coordinates.X = x
		record.Coordinates.Y = y
	} else {
		errorMsg := fmt.Sprint("failed to parse coordinates:", coordinatesStr)
		return record, errors.New(errorMsg)
	}

	record.Date = timestamp
	record.User = userStr
	record.Color = color

	return record, nil
}

// parseHexColor parses a hexadecimal color string (e.g., "#RRGGBB" or "#RRGGBBAA")
// and returns a color.Color value representing the color. It validates the input
// format, converts it to RGBA, and handles an optional alpha channel (AA).
// If the input is invalid, it returns an error.
func parseHexColor(hex string) (color.Color, error) {
	hexPattern := regexp.MustCompile(`^#([0-9A-Fa-f]{6}|[0-9A-Fa-f]{8})$`)

	if !hexPattern.MatchString(hex) {
		return nil, fmt.Errorf("invalid hex color format: %s", hex)
	}

	if hex[0] == '#' {
		hex = hex[1:]
	}

	r, _ := strconv.ParseUint(hex[0:2], 16, 8)
	g, _ := strconv.ParseUint(hex[2:4], 16, 8)
	b, _ := strconv.ParseUint(hex[4:6], 16, 8)

	var a uint64
	if len(hex) == 8 {
		a, _ = strconv.ParseUint(hex[6:8], 16, 8)
	} else {
		a = 255
	}

	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}, nil
}
