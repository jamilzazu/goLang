package rate

import "errors"

func CalculateRate(amount float64) (float64, error) {

	if amount <= 0 {
		return 0.0, errors.New("amount must be greater than 0")
	}

	if amount >= 1000 && amount < 20000 {
		return 10, nil
	}

	if amount >= 20000 {
		return 20, nil
	}
	return 5, nil
}

func CalculateRateFloat64(amount float64) float64 {

	if amount <= 0 {
		return 0.0
	}

	if amount >= 1000 && amount < 20000 {
		return 10
	}

	if amount >= 20000 {
		return 20
	}
	return 5
}

type Repository interface {
	SaveRate(amount float64) error
}

func CalculateRateAndSave(amount float64, repository Repository) error {
	rate := CalculateRateFloat64(amount)
	return repository.SaveRate(rate)
}
