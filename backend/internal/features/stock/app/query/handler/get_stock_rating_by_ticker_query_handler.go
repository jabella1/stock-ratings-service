package handler

import (
	"context"

	"github.com/jabella1/stock-ratings-service/internal/features/stock/app/dto"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/app/interfaces"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/app/query"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/domain/repositories"
)

type GetStockRatingByTickerQueryHandler struct {
	stockRatingRepository repositories.StockRatingRepository
}

func CreateGetStockRatingByTickerQueryHandler(stockRatingRepository repositories.StockRatingRepository) interfaces.GetStockRatingByTickerQueryHandler {
	return &GetStockRatingByTickerQueryHandler{
		stockRatingRepository: stockRatingRepository,
	}
}

func (s *GetStockRatingByTickerQueryHandler) GetStockRatingByTicker(context context.Context, getStockRatingByTickerQuery *query.GetStockRatingByTickerQuery) (*query.GetStockRatingByTickerResult, error) {
	stockRating, err := s.stockRatingRepository.GetStockRatingByTicker(context, getStockRatingByTickerQuery.Ticker)
	if err != nil {
		return nil, err
	}

	return &query.GetStockRatingByTickerResult{
		Result: &dto.GetStockRatingByTickerResult{
			Ticker:     stockRating.GetTicker(),
			Company:    stockRating.GetCompany(),
			Brokerage:  stockRating.GetBrokerage(),
			Action:     stockRating.GetAction(),
			RatingFrom: stockRating.GetRatingFrom(),
			RatingTo:   stockRating.GetRatingTo(),
			TargetFrom: stockRating.GetTargetFrom(),
			TargetTo:   stockRating.GetTargetTo(),
		},
	}, nil
}
