package rate

import "time"

func CalculateRate(amount float64) float64 {
	if amount == 0 {
		return 0
	}

	if amount >= 1000 {
		return 10
	}
	return 5
}

func CalculateRateTimeSleep(amount float64) float64 {
	time.Sleep(time.Millisecond)
	if amount == 0 {
		return 0
	}

	if amount >= 1000 {
		return 10
	}
	return 5
}

func CalculateRateFuzz(amount float64) float64 {

	if amount <= 0 {
		return 0
	}

	if amount >= 1000 && amount < 20000 {
		return 10
	}

	if amount >= 20000 {
		return 20
	}
	return 5
}
