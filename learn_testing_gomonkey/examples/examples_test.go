package examples

import (
	"fmt"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
)

// test without mocking with gomonkey
func TestAdd(t *testing.T) {
	result := Add(3, 4)
	expected := 7
	assert.EqualValues(t, result, expected, "The patched function did not return the expected result.")
}

// Test using gomonkey to monkey patch greet
func TestAgainAdd(t *testing.T) {
	// Monkey patch the wait function
	patches := gomonkey.ApplyFunc(wait, func() {
		fmt.Println("Hi, Welcome to mocked wait function in golang!")
		time.Sleep(1 * time.Second)
	})
	defer patches.Reset() // Ensure the patch is removed after the test

	// Test the patched behavior
	result := Add(2, 3)
	expected := 5

	assert.Equal(t, expected, result, "The patched function did not return the expected result.")
}
