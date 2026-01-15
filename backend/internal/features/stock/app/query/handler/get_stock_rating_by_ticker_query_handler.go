package handler

import (
	"github.com/jabella1/stock-ratings-service/internal/features/stock/app/dto"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/app/interfaces"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/app/query"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/domain/interface/repository"
)

type StockRatingService struct {
	stockRatingRepository repository.StockRatingRepository
}

func CreateGetStockRatingByTickerQueryHandler(stockRatingRepository repository.StockRatingRepository) interfaces.GetStockRatingByTickerQueryHandler {
	return &StockRatingService{
		stockRatingRepository: stockRatingRepository,
	}
}

func (s *StockRatingService) GetStockRatingByTicker(getStockRatingByTickerQuery *query.GetStockRatingByTickerQuery) (*query.GetStockRatingByTickerResult, error) {
	stockRating, err := s.stockRatingRepository.GetStockRatingByTicker(getStockRatingByTickerQuery.Ticker)
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
