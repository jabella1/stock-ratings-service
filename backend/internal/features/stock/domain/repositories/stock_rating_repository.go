package repositories

import (
	"github.com/jabella1/stock-ratings-service/internal/features/stock/domain/entities"
)

type StockRatingRepository interface {
	GetStockRatingByTicker(ticker string) (*entities.StockRating, error)
}
