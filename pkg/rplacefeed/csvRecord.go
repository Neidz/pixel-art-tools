package rplacefeed

import (
	"image/color"
	"time"
)

type Coordinates struct {
	X         int
	Y         int
	Rectangle Rectangle
	Circle    Circle
}

func (c *Coordinates) IsCircle() bool {
	return c.Circle != Circle{}
}

func (c *Coordinates) IsRectangle() bool {
	return c.Rectangle != Rectangle{}
}

func (c *Coordinates) IsPoint() bool {
	return c.Circle == Circle{} && c.Rectangle == Rectangle{}
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
	Color       color.Color
}
