package mypackage

// Service defines the interface for the service.
type Service interface {
	DoSomething(input string) (string, error)
}

// RealService is a real implementation of the Service.
type RealService struct{}

// DoSomething is the real implementation of the service.
func (rs *RealService) DoSomething(input string) (string, error) {
	return "Processed: " + input, nil
}
