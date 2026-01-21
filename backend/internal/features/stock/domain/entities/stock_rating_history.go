package entities

import (
	"github.com/jabella1/stock-ratings-service/internal/features/stock/domain/domainValidation"
	valueobjects "github.com/jabella1/stock-ratings-service/internal/features/stock/domain/valueObjects"
)

type StockRatingHistory struct {
	stockRatingId   int64
	oldCurrentPrice valueobjects.Price
	newCurrentPrice valueobjects.Price
	oldUpside       *float32
	newUpside       float32
}

func CreateStockRatingHistory(stockRatingId int64, oldCurrentPrice float64, newCurrentPrice float64,
	oldUpside *float32, newUpside float32) (*StockRatingHistory, error) {
	err := domainValidation.NonNegativeOrZeroId(stockRatingId, "stockRatingId")
	if err != nil {
		return nil, err
	}
	oldCurrentPriceVO, err := valueobjects.CreatePrice(&oldCurrentPrice)
	if err != nil {
		return nil, err
	}
	newCurrentPriceVO, err := valueobjects.CreatePrice(&newCurrentPrice)
	if err != nil {
		return nil, err
	}

	return &StockRatingHistory{
		stockRatingId:   stockRatingId,
		oldCurrentPrice: oldCurrentPriceVO,
		newCurrentPrice: newCurrentPriceVO,
		oldUpside:       oldUpside,
		newUpside:       newUpside,
	}, nil
}

func (stockRatingHistory *StockRatingHistory) GetStockRatingId() int64 {
	return stockRatingHistory.stockRatingId
}

func (stockRatingHistory *StockRatingHistory) GetOldCurrentPrice() float64 {
	return stockRatingHistory.oldCurrentPrice.GetValue()
}

func (stockRatingHistory *StockRatingHistory) GetNewCurrentPrice() float64 {
	return stockRatingHistory.newCurrentPrice.GetValue()
}

func (stockRatingHistory *StockRatingHistory) GetOldUpside() *float32 {
	return stockRatingHistory.oldUpside
}

func (stockRatingHistory *StockRatingHistory) GetNewUpside() float32 {
	return stockRatingHistory.newUpside
}
