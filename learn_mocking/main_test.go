package main

import (
	"errors"

	"testing"

	"example.com/u/mockdemo/mypackage"
	"example.com/u/mockdemo/mypackage/mocks"
	"github.com/golang/mock/gomock"
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

func TestService_DoSomething(t *testing.T) {
	// Create a controller for managing mock lifecycle
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock instance
	mockService := mocks.NewMockService(ctrl)

	// Define behavior of mock
	mockService.EXPECT().
		DoSomething("hello").
		Return("Mocked: hello", nil).
		Times(1)

	mockService.EXPECT().
		DoSomething("error").
		Return("", errors.New("mock error")).
		Times(1)

	// Use the mock in the test
	result, err := mockService.DoSomething("hello")
	if err != nil || result != "Mocked: hello" {
		t.Errorf("expected Mocked: hello, got %s with error %v", result, err)
	}

	result, err = mockService.DoSomething("error")
	if err == nil || err.Error() != "mock error" {
		t.Errorf("expected mock error, got %v", err)
	}
}
