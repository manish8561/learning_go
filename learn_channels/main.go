package main

import (
	"demo/examples"
	"demo/fourth"
	"demo/second"
	"demo/third"
)

func main() {

	//case 1
	examples.Example1()

	//case 2
	examples.Example2()

	//case 3
	examples.Example3()

	//case 4
	examples.Example4()

	//case 5
	examples.Example5()

	//case 6
	examples.Example6()

	//case 7
	examples.Example7()

	//case 8
	examples.Example8()

	//case 9
	examples.Example9()

	//case 10
	examples.Example10()

	//case 11
	examples.Example11()

	// synchronization example
	second.Example1()

	// memory management example
	third.Example1()

	// basic example
	examples.Example12()

	// sync map example
	fourth.Example()

}
