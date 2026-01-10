package repositories

import (
	"github.com/jabella1/stock-ratings-service/internal/features/common/pagination"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/app/query"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/domain/entities"
)

type StockRatingRepository interface {
	GetStockRatingByTicker(ticker string) (*entities.StockRating, error)
	GetListStockRating(getListStockRatingQuery *query.GetListStockRatingQuery) (*pagination.PaginatedList[entities.StockRating], error)
}
