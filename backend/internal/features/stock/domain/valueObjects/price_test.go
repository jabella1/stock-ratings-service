package valueobjects

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePrice_Negative(t *testing.T) {
	price, err := CreatePrice(ptrFloat64(-1))
	assert.Error(t, err)
	assert.Zero(t, price)
}

func TestCreatePrice_Success(t *testing.T) {
	price, err := CreatePrice(ptrFloat64(2))
	assert.NoError(t, err)
	assert.NotNil(t, price)
	assert.Equal(t, float64(2), price.GetValue())
}

func TestCreatePrice_Zero(t *testing.T) {
	price, err := CreatePrice(ptrFloat64(0))
	assert.NoError(t, err)
	assert.NotNil(t, price)
	assert.Equal(t, float64(0), price.GetValue())
}

func TestCreatePrice_Nil(t *testing.T) {
	price, err := CreatePrice(nil)
	assert.Error(t, err)
	assert.Zero(t, price)
}

func ptrFloat64(value float64) *float64 {
	return &value
}
