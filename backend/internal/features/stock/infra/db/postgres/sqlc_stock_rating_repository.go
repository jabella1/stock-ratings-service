package postgres

import (
	"context"

	"github.com/jabella1/stock-ratings-service/internal/features/common/utils"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/domain/entities"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/domain/pagination"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/domain/repositories"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/infra/db/sqlc"
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

func (sqlcRepository *SqlcStockRatingRepository) GetListStockRating(search *string, pageSize *int32, pageNumber *int32, orderBy string, orderDirection string) (*pagination.PaginatedList[entities.StockRating], error) {
	context := context.Background()
	sqlcStockRatings, err := sqlcRepository.queries.ListStockRatings(context, sqlc.ListStockRatingsParams{
		Column1: *search,
		Column2: orderBy,
		Column3: orderDirection,
		Column4: *pageSize,
		Column5: utils.CalculateOffset(*pageNumber, *pageSize),
	})

	if err != nil {
		return nil, err
	}

	totalRecords, err := sqlcRepository.queries.CountStockRatings(context,
		*search,
	)

	if err != nil {
		return nil, err
	}

	var stockRatings []entities.StockRating
	for _, sqlcStockRating := range sqlcStockRatings {
		stockRating, err := fromSqlcStockRatingToEntity(&sqlcStockRating)
		if err != nil {
			return nil, err
		}
		stockRatings = append(stockRatings, *stockRating)
	}

	return &pagination.PaginatedList[entities.StockRating]{
		Results:      &stockRatings,
		TotalRecords: totalRecords,
	}, nil
}

func fromSqlcStockRatingToEntity(sqlcStockRating *sqlc.ChallengeStockRating) (*entities.StockRating, error) {
	return entities.CreateStockRating(
		sqlcStockRating.Ticker,
		sqlcStockRating.Company,
		&sqlcStockRating.Brokerage.String,
		&sqlcStockRating.Action.String,
		&sqlcStockRating.RatingFrom.String,
		&sqlcStockRating.RatingTo.String,
		utils.Float64FromNumeric(sqlcStockRating.TargetFrom),
		utils.Float64FromNumeric(sqlcStockRating.TargetTo),
		*utils.Float64FromNumeric(sqlcStockRating.Upside),
		*utils.Float64FromNumeric(sqlcStockRating.ChangeTarget),
		*utils.Float64FromNumeric(sqlcStockRating.CurrentPrice),
	)
}
