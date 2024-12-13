package main

import (
	"errors"

	"testing"

	"example.com/u/mockdemo/mypackage"
)

func TestMockService(t *testing.T) {
	// Create a mock service
	mockService := &mypackage.MockService{
		DoSomethingFunc: func(input string) (string, error) {
			if input == "error" {
				return "", errors.New("mock error")
			}
			return "Mocked: " + input, nil
		},
	}

	// Test success case
	result, err := mockService.DoSomething("test")
	if err != nil || result != "Mocked: test" {
		t.Errorf("expected Mocked: test, got %s with error %v", result, err)
	}

	// Test error case
	_, err = mockService.DoSomething("error")
	if err == nil || err.Error() != "mock error" {
		t.Errorf("expected mock error, got %v", err)
	}
}
