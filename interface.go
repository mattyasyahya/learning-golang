package main

import (
	"fmt"
	"math"
)

// Shape interface
type Shape interface {
	area() float64
}

// Rectangle is a rectangle data type
type Rectangle struct {
	x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

// Circle is a circle data type
type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {
	circle := Circle{10, 10, 10}
	rectangle := Rectangle{0, 0, 10, 10}

	fmt.Println(getArea(&circle))
	fmt.Println(getArea(&rectangle))
}
