package repositories

import (
	"github.com/jabella1/stock-ratings-service/internal/features/stock/domain/entities"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/domain/pagination"
)

type StockRatingRepository interface {
	GetStockRatingByTicker(ticker string) (*entities.StockRating, error)
	GetListStockRating(search *string, pageSize *int32, pageNumber *int32, orderBy string, orderDirection string) (*pagination.PaginatedList[entities.StockRating], error)
}
