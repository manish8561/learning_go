package examples

import (
	"fmt"
	"sync"
)
// singleton example with sync.once
// Singleton struct - the single instance we want to create.
type Singleton struct {
	value int
}

// The instance of the Singleton. It's initialized lazily.
var instance *Singleton
var once sync.Once // Ensures that the instance is created only once

// GetInstance function returns the singleton instance.
func GetInstance() *Singleton {
	once.Do(func() {
		fmt.Println("Creating a new instance.")
		instance = &Singleton{value: 42}
	})
	return instance
}

func Example1() {
	// Get the first instance
	s1 := GetInstance()
	fmt.Println("Instance 1 value:", s1.value)

	// Get the second instance (this won't create a new one)
	s2 := GetInstance()
	fmt.Println("Instance 2 value:", s2.value)

	// Both instances should point to the same address
	if s1 == s2 {
		fmt.Println("Both instances are the same.")
	}
}
