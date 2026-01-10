package query

import "github.com/jabella1/stock-ratings-service/internal/features/stock/app/dto"

type GetStockRatingByTickerQuery struct {
	Ticker string
}

type GetStockRatingByTickerResult struct {
	Result *dto.GetStockRatingByTickerResult
}
