package repositories

import (
	"github.com/jabella1/stock-ratings-service/internal/app/features/stock/domain/entities"
)

type StockRatingRepository interface {
	GetStockRatingByTicker(ticker string) (*entities.StockRating, error)
}
