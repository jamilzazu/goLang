package rate

import (
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
