package main

import (
	"fmt"
	"os"
	"text/template"
)

type Student struct {
	Name  string
	Marks int
	Id    string
}

func main() {
	std1 := Student{"Manish Sharma", 95, "1"}

	// tex templating
	tmp1 := template.New("test_1")

	tmp1, _ = tmp1.Parse("Hello {{.Name}}, your marks are {{.Marks}}%!\n")

	err := tmp1.Execute(os.Stdout, std1)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("-----------------------------------------")
	// html templating
	t, err := template.ParseFiles("index.html")

	if err != nil {
		fmt.Println(err)
	}

	err = t.Execute(os.Stdout, std1)

	if err != nil {
		fmt.Println(err)
	}
}
