package main

// Structural Design Pattern: Adaptor Method

import "fmt"

// Bird interface defines the behavior of a bird.
type Bird interface {
	MakeSound() string
}

// Sparrow is a concrete implementation of Bird.
type Sparrow struct{}

func (s Sparrow) MakeSound() string {
	return "Chirp Chirp"
}

// Toy interface defines the behavior of a toy.
type Toy interface {
	Squeak() string
}

// ToyDuck is a concrete implementation of Toy.
type ToyDuck struct{}

func (td ToyDuck) Squeak() string {
	return "Squeak!"
}

// BirdAdapter adapts a Bird to be used as a Toy.
type BirdAdapter struct {
	Bird Bird
}

// Squeak implements the Toy interface, adapting the bird's sound.
func (ba BirdAdapter) Squeak() string {
	return ba.Bird.MakeSound() // Adapting the bird's sound as the toy's squeak
}

func main() {
	// Create a Sparrow instance
	sparrow := Sparrow{}

	// Create a ToyDuck instance
	toyDuck := ToyDuck{}

	// Create a BirdAdapter instance, adapting the Sparrow to the Toy interface
	birdAdapter := BirdAdapter{Bird: sparrow}

	// Use the toy duck
	fmt.Println("ToyDuck:", toyDuck.Squeak())

	// Use the bird adapter to make the sparrow act like a toy duck
	fmt.Println("BirdAdapter (Sparrow as Toy):", birdAdapter.Squeak())
}