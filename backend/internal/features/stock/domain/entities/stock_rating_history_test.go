package entities_test

import (
	"testing"

	"github.com/jabella1/stock-ratings-service/internal/features/stock/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestCreateStockRatingHistory_Success(t *testing.T) {
	oldPrice := 100.0
	newPrice := 120.0
	oldUpside := float32(5.5)
	newUpside := float32(10.0)

	stockRatingHistory, err := entities.CreateStockRatingHistory(1, oldPrice, newPrice, &oldUpside, newUpside)
	assert.NoError(t, err)
	assert.NotNil(t, stockRatingHistory)

	assert.Equal(t, int64(1), stockRatingHistory.GetStockRatingId())
	assert.Equal(t, oldPrice, stockRatingHistory.GetOldCurrentPrice())
	assert.Equal(t, newPrice, stockRatingHistory.GetNewCurrentPrice())
	assert.Equal(t, &oldUpside, stockRatingHistory.GetOldUpside())
	assert.Equal(t, newUpside, stockRatingHistory.GetNewUpside())
}

func TestCreateStockRatingHistory_NegativeID(t *testing.T) {
	oldPrice := 100.0
	newPrice := 120.0
	oldUpside := float32(5.5)
	newUpside := float32(10.0)

	stockRatingHistory, err := entities.CreateStockRatingHistory(-1, oldPrice, newPrice, &oldUpside, newUpside)
	assert.Error(t, err)
	assert.Nil(t, stockRatingHistory)
}

func TestCreateStockRatingHistory_NilOldUpside(t *testing.T) {
	oldPrice := 50.0
	newPrice := 75.0
	newUpside := float32(12.5)

	stockRatingHistory, err := entities.CreateStockRatingHistory(10, oldPrice, newPrice, nil, newUpside)
	assert.NoError(t, err)
	assert.NotNil(t, stockRatingHistory)

	assert.Nil(t, stockRatingHistory.GetOldUpside())
	assert.Equal(t, newUpside, stockRatingHistory.GetNewUpside())
}
