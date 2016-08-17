package main

import (
	"fmt"
	"math"
)

// Rectangle is a rectangle data type
type Rectangle struct {
	x1, y1, x2, y2 float64
}

// Circle is a circle data type
type Circle struct {
	x, y, r float64
}

// Name is person name data type
type Name struct {
	firstName, lastName string
}

// Person is person data type
type Person struct {
	name    Name
	address string
}

func changePersonName(person *Person) {
	person.name.firstName = "Ganti"
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
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

func main() {
	fmt.Println("Start program")

	circle := Circle{
		x: 10,
		y: 10,
		r: 10,
	}

	circle.r = 100

	fmt.Println(circle)

	person := Person{
		name: Name{
			firstName: "Eko Kurniawan",
			lastName:  "Khannedy",
		},
		address: "Subang",
	}
	fmt.Println(person)

	changePersonName(&person)
	fmt.Println(person)

	fmt.Println(circle.area())

	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())

	fmt.Println("End program")
}
