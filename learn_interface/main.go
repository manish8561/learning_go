package main

import (
	"fmt"
	"math"

	"example.com/demo/examples"
	"example.com/demo/second"
)

// common function
type Shape interface {
	Area() float32
}

type Square struct {
	length float32
}

// implementation of interface with struct
func (s *Square) Area() float32 {
	return s.length * s.length
}

type Rectangle struct {
	length float32
	width  float32
}

// implementation of interface with struct
func (s *Rectangle) Area() float32 {
	return s.length * s.width
}

type Circle struct {
	radius float32
}

// implementation of interface with struct
func (c *Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

func getArea(s Shape) float32 {
	return s.Area()
}

func main() {
	s := Square{length: 2.5}
	c := Circle{radius: 2.5}
	r := Rectangle{length: 4.5, width: 2.5}

	// make slice for common types using interface
	shapes := []Shape{&s, &c, &r}

	//range on slice
	for i, v := range shapes {
		fmt.Println(i, " Area: ", getArea(v))
	}

	// any type as interface
	var test interface{}
	test = "some text"
	test = 10

	val, ok := test.(int)
	if ok {
		fmt.Println(val)
	}

	//Second Example
	examples.Example()

	//second Example
	second.Example()
}
