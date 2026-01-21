package handler

import (
	"context"
	"errors"
	"testing"

	"github.com/jabella1/stock-ratings-service/internal/features/common/utils"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/app/command"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/domain/repositories"
	"github.com/jabella1/stock-ratings-service/internal/mocks"
	"github.com/stretchr/testify/mock"
)

func TestSaveStockRatingCommandHandler(t *testing.T) {
	dbErr := errors.New("database crash")
	context := context.Background()
	saveStockRatingCommand := &command.SaveStockRatingCommand{
		Ticker:       "JFAM",
		Company:      "JuanFCompany",
		Brokerage:    nil,
		Action:       "target lowered by",
		RatingFrom:   "buy",
		RatingTo:     "buy",
		TargetFrom:   8,
		TargetTo:     6,
		CurrentPrice: 1.34,
	}
	var upside = utils.CalculatePercentageChange(&saveStockRatingCommand.CurrentPrice, &saveStockRatingCommand.TargetTo)

	var tests = []struct {
		name        string
		changePrice bool
		shouldFail  bool
	}{
		{"succeed save with history, changeprice", true, true},
		{"succeed save without history, no changeprice", false, true},
		{"succeed rollback", true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStockRatingRepository := mocks.NewStockRatingRepository(t)
			mockUnitOfWork := mocks.NewUnitOfWork(t)
			mockTransaction := mocks.NewDatabaseTransaction(t)
			saveStockRatingCommandHandler := CreateSaveStockRatingCommandHandler(mockStockRatingRepository, mockUnitOfWork)

			mockUnitOfWork.On("BeginTransaction", context).Return(mockTransaction, nil)
			mockTransaction.On("Rollback", context).Return(nil)

			if !tt.shouldFail {
				mockStockRatingRepository.On("SaveStockRating", context, mock.Anything, mock.Anything).Return(&repositories.UpsertResult{
					ID:            1,
					PriceChanged:  tt.changePrice,
					UpsideChanged: tt.changePrice,
					OldPrice:      1.34,
					OldUpside:     3.4776,
				}, nil)

				if tt.changePrice {
					mockStockRatingRepository.On("SaveStockRatingHistory", context, mock.Anything, mock.Anything).Return(nil)
				}
				mockTransaction.On("Commit", context).Return(nil)

			} else {
				mockStockRatingRepository.On("SaveStockRating", context, mock.Anything, mock.Anything).Return(nil, dbErr)
				mockTransaction.AssertNotCalled(t, "Commit", context)
			}

			response, err := saveStockRatingCommandHandler.SaveStockRating(context, saveStockRatingCommand)

			if tt.shouldFail {
				if err == nil {
					t.Errorf("expected error but got nil")
				}
			} else {

				if err != nil {
					t.Errorf("expected no error, but got %v", err)
				}
				if response.Data.Upside != float32(upside) {
					t.Errorf("expected upside= %v got=%v", upside, response.Data.Upside)
				}
			}

			mockStockRatingRepository.AssertExpectations(t)
			mockUnitOfWork.AssertExpectations(t)
			mockTransaction.AssertExpectations(t)
		})
	}
}
