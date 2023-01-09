package rate

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateRate(t *testing.T) {
	rate, err := CalculateRate(1000.0)
	assert.Nil(t, err)
	assert.Equal(t, 10.0, rate)

	rate, err = CalculateRate(0)

	assert.Error(t, err, "amount must be greater than 0")
	assert.Equal(t, 0.0, rate)
	assert.Contains(t, err.Error(), "greater than 0")
}

func TestCalculateRateAndSave(t *testing.T) {
	repository := &RateRepositoryMock{}
	repository.On("SaveRate", 10.0).Return(nil).Once()
	repository.On("SaveRate", 0.0).Return(errors.New("error saving rate"))

	err := CalculateRateAndSave(1000.0, repository)
	assert.Nil(t, err)

	err = CalculateRateAndSave(0.0, repository)
	assert.Error(t, err, "error saving rate")

	repository.AssertExpectations(t)
	repository.AssertNumberOfCalls(t, "SaveRate", 2)
}
