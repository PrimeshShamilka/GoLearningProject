package oop

import (
	"fmt"
	"math"
)

type Square struct {
	Length float64
}

type Circle struct {
	Radius float64
}

// Shape is an interface
// method area returns the area of the shape in float64
type Shape interface {
	Area() float64
}

func (s Square) Area() float64 {
	return s.Length * s.Length
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func sumAreas(shapes []Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

// ignore redeclare of main error for learning purpose
func main() {
	s := Square{Length: 5}
	c := Circle{Radius: 3}
	shapes := []Shape{s, c}
	fmt.Println("Total Area:", sumAreas(shapes))
}
