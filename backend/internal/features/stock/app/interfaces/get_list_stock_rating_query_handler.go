package interfaces

import (
	"context"

	"github.com/jabella1/stock-ratings-service/internal/features/stock/app/query"
)

type GetListStockRatingQueryHandler interface {
	GetListStockRating(context context.Context, getListStockRatingQuery *query.GetListStockRatingQuery) (*query.GetListStockRatingResult, error)
}
