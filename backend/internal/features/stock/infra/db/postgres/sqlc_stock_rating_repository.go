package postgres

import (
	"context"

	"github.com/jabella1/stock-ratings-service/internal/features/common/utils"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/app/query"
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

func (r *SqlcStockRatingRepository) GetListStockRating(getListStockRatingQuery *query.GetListStockRatingQuery) (*pagination.PaginatedList[entities.StockRating], error) {
	context := context.Background()
	sqlcStockRatings, err := r.queries.ListStockRatings(context, sqlc.ListStockRatingsParams{
		Column1: *getListStockRatingQuery.Search,
		Column2: *getListStockRatingQuery.OrderBy,
		Column3: *getListStockRatingQuery.OrderDirection,
		Column4: *getListStockRatingQuery.PageSize,
		Column5: utils.CalculateOffset(*getListStockRatingQuery.PageNumber, *getListStockRatingQuery.PageSize),
	})
	if err != nil {
		return nil, err
	}

	totalRecords, err := r.queries.CountStockRatings(context,
		*getListStockRatingQuery.Search,
	)

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
	)
}
