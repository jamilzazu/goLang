package rate

import "testing"

// Commands go test . go test -v

func TestCalculateRate(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateRate(amount)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}
