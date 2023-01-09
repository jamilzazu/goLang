package rate

import "testing"

// Commands
// go test .
// go test -v
// go test -coverprofile=coverage.out && go tool cover -html=coverage.out

func TestCalculateRate(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateRate(amount)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestCalculateRateBatch(t *testing.T) {

	type calcRate struct {
		amount, expect float64
	}

	table := []calcRate{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
		//{0.0, 0.0},
	}

	for _, item := range table {
		result := CalculateRate(item.amount)

		if result != item.expect {
			t.Errorf("Expected %f but got %f", item.expect, result)
		}
	}
}
