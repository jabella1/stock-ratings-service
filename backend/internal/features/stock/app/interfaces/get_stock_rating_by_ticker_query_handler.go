package interfaces

import "github.com/jabella1/stock-ratings-service/internal/app/features/stock/app/query"

type GetStockRatingByTickerQueryHandler interface {
	GetStockRatingByTicker(getStockRatingByTickerQuery *query.GetStockRatingByTickerQuery) (*query.GetStockRatingByTickerResult, error)
}
