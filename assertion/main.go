// type-assertion let's you determine and extract an interface value and its type
package main

import (
	"fmt"
	"math"
) //later usage for pie

// define a interface
type Shape interface {
	Area() float64
}

// implementing the interface of a Rectangle struct
type Rectangle struct {
	length, width float64
}

// implementing the interface of a Circle struct
type Circle struct {
	radius float64
}

// Defining areas for both shapes
// Rectangle
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// circle
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func identify(shape Shape) {
	switch s := shape.(type) {
	case Rectangle:
		fmt.Printf("Area of Rectangle is %.2f\n", s.Area())
	case Circle:
		fmt.Printf("Area of Circle is %.2f\n", s.Area())
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	r := Rectangle{length: 10, width: 5}
	c := Circle{radius: 5}

	identify(r)
	identify(c)
}
