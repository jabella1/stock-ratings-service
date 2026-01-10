package interfaces

import "github.com/jabella1/stock-ratings-service/internal/features/stock/app/query"

type GetListStockRatingQueryHandler interface {
	GetListStockRating(getListStockRatingQuery *query.GetListStockRatingQuery) (*query.GetListStockRatingResult, error)
}
