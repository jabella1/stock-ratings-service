package interfaces

import (
	"context"

	"github.com/jabella1/stock-ratings-service/internal/features/stock/app/query"
)

type GetStockRatingByTickerQueryHandler interface {
	GetStockRatingByTicker(context context.Context, getStockRatingByTickerQuery *query.GetStockRatingByTickerQuery) (*query.GetStockRatingByTickerResult, error)
}
