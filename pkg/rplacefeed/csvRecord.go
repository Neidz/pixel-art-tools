package rplacefeed

import "time"

type Coordinates struct {
	X         int
	Y         int
	Rectangle Rectangle
	Circle    Circle
}

type Rectangle struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}

type Circle struct {
	X int
	Y int
	R int
}

type CSVRecord struct {
	Date        time.Time
	User        string
	Coordinates Coordinates
	Color       string
}
