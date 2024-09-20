package math

import "testing"

func TestAdd(t *testing.T) {
	actualOutput := Add(2,2)
	expectedOutput := 4

	if actualOutput != expectedOutput {
		t.Errorf("Expected Output (%d) is not equal to actual output (%d)", expectedOutput, actualOutput)
	}
}
