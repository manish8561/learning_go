package main

import (
	"fmt"

	"example.com/u/mockdemo/mypackage"
)

func main() {
	ob := mypackage.RealService{}

	result, err := ob.DoSomething("Hello World")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("mypackage.RealService: %v\n", result)
}
