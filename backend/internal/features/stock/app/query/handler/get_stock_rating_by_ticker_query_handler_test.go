package handler

import (
	"context"
	"errors"
	"testing"

	"github.com/jabella1/stock-ratings-service/internal/features/common/utils"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/app/query"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/domain/entities"
	"github.com/jabella1/stock-ratings-service/internal/mocks"
	"github.com/stretchr/testify/mock"
)

func TestGetStockRatingByTickerQueryHandler(t *testing.T) {
	dbErr := errors.New("ticker no found")
	tests := []struct {
		name        string
		ticker      string
		foundTicker bool
	}{
		{"exist ticker", "JFAM", true},
		{"no exist ticker", "JFAMNEW", false},
	}
	context := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStockRatingRepository := mocks.NewStockRatingRepository(t)
			getStockRatingByTickerQueryHandler := CreateGetStockRatingByTickerQueryHandler(mockStockRatingRepository)
			getStockRatingByTickerQuery := query.GetStockRatingByTickerQuery{Ticker: tt.ticker}
			if tt.foundTicker {
				stockRatingEntity, err := entities.CreateStockRating(tt.ticker, "JuanFCompany",
					utils.PtrString(""), utils.PtrString("target lowered by"), utils.PtrString("buy"), utils.PtrString("buy"), utils.PtrFloat64(8), utils.PtrFloat64(6), 3.4776, 1, 1.34)
				if err != nil {
					t.Fatalf("failed creating entity")
				}
				mockStockRatingRepository.On("GetStockRatingByTicker", mock.Anything, mock.Anything).Return(stockRatingEntity, nil)
			} else {
				mockStockRatingRepository.On("GetStockRatingByTicker", mock.Anything, mock.Anything).Return(nil, dbErr)
			}
			_, err := getStockRatingByTickerQueryHandler.GetStockRatingByTicker(context, &getStockRatingByTickerQuery)

			if tt.foundTicker {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
				}
			} else {
				if err == nil {
					t.Error("expected error but got nil")
				}
			}
			mockStockRatingRepository.AssertExpectations(t)
		})
	}
}
