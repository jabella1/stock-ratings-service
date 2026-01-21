package entities_test

import (
	"testing"

	"github.com/jabella1/stock-ratings-service/internal/features/stock/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestCreateStockRating_Success(t *testing.T) {
	brokerage := "Broker A"
	action := "target lowered by"
	ratingFrom := "Buy"
	ratingTo := "Buy"
	targetFrom := 3.0000
	targetTo := 2.5000
	currentPrice := 1.2600
	upside := float32(0.9841)
	changeTarget := -0.1667

	stockRating, err := entities.CreateStockRating(
		"TSLA",
		"Tesla Inc",
		&brokerage,
		&action,
		&ratingFrom,
		&ratingTo,
		&targetFrom,
		&targetTo,
		upside,
		changeTarget,
		currentPrice,
	)

	assert.NoError(t, err)
	assert.NotNil(t, stockRating)

	assert.Equal(t, "TSLA", stockRating.GetTicker())
	assert.Equal(t, "Tesla Inc", stockRating.GetCompany())
	assert.Equal(t, &brokerage, stockRating.GetBrokerage())
	assert.Equal(t, &action, stockRating.GetAction())
	assert.Equal(t, &ratingFrom, stockRating.GetRatingFrom())
	assert.Equal(t, &ratingTo, stockRating.GetRatingTo())
	assert.Equal(t, targetFrom, stockRating.GetTargetFrom())
	assert.Equal(t, targetTo, stockRating.GetTargetTo())
	assert.Equal(t, upside, stockRating.GetUpside())
	assert.Equal(t, changeTarget, stockRating.GetChangeTarget())
	assert.Equal(t, currentPrice, stockRating.GetCurrentPrice())
}

func TestCreateStockRating_EmptyTicker(t *testing.T) {
	_, err := entities.CreateStockRating(
		"",
		"Tesla Inc",
		nil, nil, nil, nil,
		nil, nil,
		.9841, -0.1667, 1.2600,
	)
	assert.Error(t, err)
}

func TestCreateStockRating_EmptyCompany(t *testing.T) {
	_, err := entities.CreateStockRating(
		"TSLA",
		"",
		nil, nil, nil, nil,
		nil, nil,
		.9841, -0.1667, 1.2600,
	)
	assert.Error(t, err)
}

func TestCreateStockRating_NilOptionalFields(t *testing.T) {
	targetFrom := 3.0000
	targetTo := 2.50000
	currentPrice := 1.2600

	stockRating, err := entities.CreateStockRating(
		"AAPL",
		"Apple Inc",
		nil, nil, nil, nil,
		&targetFrom,
		&targetTo,
		5.0,
		15.0,
		currentPrice,
	)

	assert.NoError(t, err)
	assert.NotNil(t, stockRating)
	assert.Nil(t, stockRating.GetBrokerage())
	assert.Nil(t, stockRating.GetAction())
	assert.Nil(t, stockRating.GetRatingFrom())
	assert.Nil(t, stockRating.GetRatingTo())
	assert.Equal(t, targetFrom, stockRating.GetTargetFrom())
	assert.Equal(t, targetTo, stockRating.GetTargetTo())
}
