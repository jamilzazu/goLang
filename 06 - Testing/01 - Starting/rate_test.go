package rate

import "testing"

// Commands
// go test .
// go test -v

// #Coverage
// go test -coverprofile=coverage.out && go tool cover -html=coverage.out

// #Benchmark
// go test -bench .
// go test -bench . -run=^#
// go test -bench . -run=^# -count=10
// go test -bench . -run=^# -count=10 -benchtime=3s
// go test -bench . -run=^# -benchmem

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

func BenchmarkCalculateRateBatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateRate(500.0)
	}
}

func BenchmarkCalculateRateBatchTimeSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateRateTimeSleep(500.0)
	}
}
