package main

import "fmt"

// Animal interface defines methods that all concrete animals should implement.
type Animal interface {
	Speak() string
}

// Dog is a concrete type that implements the Animal interface.
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

// Cat is a concrete type that implements the Animal interface.
type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

// Cow is a concrete type that implements the Animal interface.
type Cow struct{}

func (c Cow) Speak() string {
	return "Moo!"
}

// AnimalFactory is the factory function that creates animals.
func AnimalFactory(animalType string) (Animal, error) {
	switch animalType {
	case "dog":
		return &Dog{}, nil
	case "cat":
		return &Cat{}, nil
	case "cow":
		return &Cow{}, nil
	default:
		return nil, fmt.Errorf("unknown animal type: %s", animalType) // creating the error here
	}
}

func main() {
	// Create a Dog
	dog, _ := AnimalFactory("dog")
	fmt.Println("Dog says:", dog.Speak())

	// Create a Cat
	cat, _ := AnimalFactory("cat")
	fmt.Println("Cat says:", cat.Speak())

	// Create a Cow
	cow, _ := AnimalFactory("cow")
	fmt.Println("Cow says:", cow.Speak())

	// Try to create an unknown animal type
	unknownAnimal, err := AnimalFactory("lion")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Unknown animal says:", unknownAnimal.Speak())
	}
}
