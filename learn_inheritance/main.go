package main

import "fmt"

// interview question
type Person struct{
	Name string
	Age int
}

func (p *Person) Intro(){
	fmt.Println(p.Name, p.Age)
}

type Employee struct {
	*Person //Embedding
	EmployeeId int
}

func (e *Employee) Intro(){
	fmt.Println(e.Name, e.Age, e.EmployeeId)
}

func main(){
	fmt.Println("Inheritance")

	p := Person{Name: "Manish Sharma", Age:38}

	e := Employee{Person: &p, EmployeeId: 1}

	e.Intro()
	// p.Intro()

}