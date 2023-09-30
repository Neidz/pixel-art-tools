package rplacefeed

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func ParseRecord(dateStr string, userStr, coordinatesStr string, colorStr string) (CSVRecord, error) {
	var record CSVRecord

	timestamp, err := time.Parse("2006-01-02 15:04:05.999 UTC", dateStr)

	if err != nil {
		return CSVRecord{}, err
	}

	rectanglePattern := regexp.MustCompile(`(\d+),(\d+),(\d+),(\d+)`)
	circlePattern := regexp.MustCompile(`{X: (\d+), Y: (\d+), R: (\d+)}`)
	simplePattern := regexp.MustCompile(`(\d+),(\d+)`)

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
		fmt.Println("failed to parse coordinates:", coordinatesStr)
		errorMsg := fmt.Sprint("failed to parse coordinates:", coordinatesStr)
		return record, errors.New(errorMsg)
	}

	record.Date = timestamp
	record.User = userStr

	record.Color = colorStr

	return record, nil
}
