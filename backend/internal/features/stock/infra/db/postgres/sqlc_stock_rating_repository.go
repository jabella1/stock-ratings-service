package postgres

import (
	"context"

	"github.com/jabella1/stock-ratings-service/internal/app/features/stock/domain/entities"
	"github.com/jabella1/stock-ratings-service/internal/app/features/stock/domain/repositories"
	"github.com/jabella1/stock-ratings-service/internal/app/features/stock/infra/db/sqlc"
)

type SqlcStockRatingRepository struct {
	queries *sqlc.Queries
}

func CreateSqlcStockRatingRepository(queries *sqlc.Queries) repositories.StockRatingRepository {
	return &SqlcStockRatingRepository{queries: queries}
}

func (r *SqlcStockRatingRepository) GetStockRatingByTicker(ticker string) (*entities.StockRating, error) {
	ctx := context.Background()
	stockRating, err := r.queries.GetStockRatingByTicker(ctx, ticker)
	if err != nil {
		return nil, err
	}
	return fromSqlcStockRatingToEntity(&stockRating)
}

func fromSqlcStockRatingToEntity(sqlcStockRating *sqlc.ChallengeStockRating) (*entities.StockRating, error) {
	return entities.CreateStockRating(
		sqlcStockRating.Ticker,
		sqlcStockRating.Company,
		&sqlcStockRating.Brokerage.String,
		&sqlcStockRating.Action.String,
		&sqlcStockRating.RatingFrom.String,
		&sqlcStockRating.RatingTo.String,
		Float64FromNumeric(sqlcStockRating.TargetFrom),
		Float64FromNumeric(sqlcStockRating.TargetTo),
	)
}
