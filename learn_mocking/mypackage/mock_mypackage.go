package mypackage

// MockService is a mock implementation of the Service interface.
type MockService struct {
	DoSomethingFunc func(input string) (string, error)
}

// DoSomething calls the mock function.
func (ms *MockService) DoSomething(input string) (string, error) {
	return ms.DoSomethingFunc(input)
}
