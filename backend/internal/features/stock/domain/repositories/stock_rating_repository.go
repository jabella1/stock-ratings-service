package repositories

import (
	"context"

	"github.com/jabella1/stock-ratings-service/internal/features/stock/domain/entities"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/domain/pagination"
	unitofwork "github.com/jabella1/stock-ratings-service/internal/features/stock/domain/unitOfWork"
)

type StockRatingRepository interface {
	GetStockRatingByTicker(context context.Context, ticker string) (*entities.StockRating, error)
	GetListStockRating(context context.Context, search *string, pageSize *int32, pageNumber *int32, orderBy string, orderDirection string,
		minUpside float32, minPrice, maxPrice float64) (*pagination.PaginatedList[entities.StockRating], error)
	SaveStockRating(context context.Context, databaseTransaction unitofwork.DatabaseTransaction, stockRating *entities.StockRating) (*UpsertResult, error)
	SaveStockRatingHistory(context context.Context, databaseTransaction unitofwork.DatabaseTransaction, stockRatingHistory *entities.StockRatingHistory) error
}

type UpsertResult struct {
	ID            int64
	PriceChanged  bool
	UpsideChanged bool
	OldPrice      float64
	OldUpside     float32
}
