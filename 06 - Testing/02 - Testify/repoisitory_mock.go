package rate

import (
	"github.com/stretchr/testify/mock"
)

type RateRepositoryMock struct {
	mock.Mock
}

func (rateRepositoryMock *RateRepositoryMock) SaveRate(rate float64) error {
	args := rateRepositoryMock.Called(rate)
	return args.Error(0)
}
