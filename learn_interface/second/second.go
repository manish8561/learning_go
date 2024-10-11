package second

import "fmt"

// interview question
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

type Square struct {
	Length float64
}

func (s Square) Area() float64 {
	return s.Length * s.Length
}

func PrintArea(s Shape) {
	fmt.Println(s.Area())
}

func Example() {
	// polymorphism done using interfaces
	r := Rectangle{10.5, 20.8}
	s := Square{2.5}

	PrintArea(r)
	PrintArea(s)
}
