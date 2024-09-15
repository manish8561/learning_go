package main
// example of generic in golang need to study more for worker
import (
	"fmt"
)

type Person interface {
	Work()
}
//generic 
type worker string
type worker2 int

func (w worker) Work() {
	fmt.Printf("%s worker1 type is working\n", w)
}
func (w worker2) Work() {
	fmt.Printf("%d worker2 type is working\n", w)
}
func DoWork[T Person](things []T) {
	for _, v := range things {
		v.Work()
	}
}

func main() {
	var a, b, c worker
	var d, e, f worker2
	a = "A"
	b = "B"
	c = "C"
	DoWork([]worker{a, b, c})
	d = 1
	e = 2
	f = 3
	DoWork([]worker2{d, e, f})
}
