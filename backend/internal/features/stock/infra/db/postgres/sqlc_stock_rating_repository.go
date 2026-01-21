package postgres

import (
	"context"
	"math"

	"github.com/jabella1/stock-ratings-service/internal/features/common/utils"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/domain/entities"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/domain/pagination"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/domain/repositories"
	unitofwork "github.com/jabella1/stock-ratings-service/internal/features/stock/domain/unitOfWork"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/infra/db/sqlc"
)

type SqlcStockRatingRepository struct {
	queries *sqlc.Queries
}

func CreateSqlcStockRatingRepository(queries *sqlc.Queries) repositories.StockRatingRepository {
	return &SqlcStockRatingRepository{queries: queries}
}

func (r *SqlcStockRatingRepository) GetStockRatingByTicker(context context.Context, ticker string) (*entities.StockRating, error) {
	stockRating, err := r.queries.GetStockRatingByTicker(context, ticker)
	if err != nil {
		return nil, err
	}
	return fromSqlcStockRatingToEntity(&stockRating)
}

func (sqlcRepository *SqlcStockRatingRepository) GetListStockRating(context context.Context, search *string, pageSize *int32, pageNumber *int32, orderBy string, orderDirection string,
	minUpside float32, minPrice, maxPrice float64) (*pagination.PaginatedList[entities.StockRating], error) {
	sqlcStockRatings, err := sqlcRepository.queries.ListStockRatings(context, sqlc.ListStockRatingsParams{
		Column1: *search,
		Column2: orderBy,
		Column3: orderDirection,
		Column4: *pageSize,
		Column5: utils.CalculateOffset(*pageNumber, *pageSize),
		Column6: utils.NumericFromFloat64(float64(minUpside)),
		Column7: utils.NumericFromFloat64(minPrice),
		Column8: utils.NumericFromFloat64(maxPrice),
	})

	if err != nil {
		return nil, err
	}

	totalRecords, err := sqlcRepository.queries.CountStockRatings(context, sqlc.CountStockRatingsParams{
		Column1: *search,
		Column2: utils.NumericFromFloat64(float64(minUpside)),
		Column3: utils.NumericFromFloat64(minPrice),
		Column4: utils.NumericFromFloat64(maxPrice),
	})

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

func (sqlcRepository *SqlcStockRatingRepository) SaveStockRating(context context.Context, databaseTransaction unitofwork.DatabaseTransaction,
	stockRating *entities.StockRating) (*repositories.UpsertResult, error) {
	stockRatingEntity, err := databaseTransaction.Queries().GetStockRatingByTicker(context, stockRating.GetTicker())
	var oldPrice float64 = 0
	var oldUpside float32 = 0
	if err == nil {
		oldPrice = *utils.Float64FromNumeric(stockRatingEntity.CurrentPrice)
		oldUpside = float32(*utils.Float64FromNumeric(stockRatingEntity.Upside))

	}

	id, err := databaseTransaction.Queries().UpsertStockRating(context, sqlc.UpsertStockRatingParams{
		Ticker:       stockRating.GetTicker(),
		Company:      stockRating.GetCompany(),
		Brokerage:    utils.TextFromString(*stockRating.GetBrokerage()),
		Action:       utils.TextFromString(*stockRating.GetAction()),
		RatingFrom:   utils.TextFromString(*stockRating.GetRatingFrom()),
		RatingTo:     utils.TextFromString(*stockRating.GetRatingTo()),
		TargetFrom:   utils.NumericFromFloat64(stockRating.GetTargetFrom()),
		TargetTo:     utils.NumericFromFloat64(stockRating.GetTargetTo()),
		Upside:       utils.NumericFromFloat64(float64(stockRating.GetUpside())),
		ChangeTarget: utils.NumericFromFloat64(stockRating.GetChangeTarget()),
		CurrentPrice: utils.NumericFromFloat64(stockRating.GetCurrentPrice()),
	})

	if err != nil {
		return nil, err
	}
	priceChanged := isValueChanged(oldPrice, stockRating.GetCurrentPrice())
	upsideChanged := isValueChanged(float64(oldUpside), float64(stockRating.GetUpside()))
	return &repositories.UpsertResult{ID: id, PriceChanged: priceChanged,
		UpsideChanged: upsideChanged, OldPrice: oldPrice, OldUpside: oldUpside}, nil
}

func isValueChanged(oldValue, newValue float64) bool {
	return math.Round(oldValue*10000) != math.Round(newValue*10000)
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
		float32(*utils.Float64FromNumeric(sqlcStockRating.Upside)),
		*utils.Float64FromNumeric(sqlcStockRating.ChangeTarget),
		*utils.Float64FromNumeric(sqlcStockRating.CurrentPrice),
	)
}

func (sqlcRepository *SqlcStockRatingRepository) SaveStockRatingHistory(context context.Context, databaseTransaction unitofwork.DatabaseTransaction,
	stockRatingHistory *entities.StockRatingHistory) error {
	err := databaseTransaction.Queries().InsertStockRatingHistory(context, sqlc.InsertStockRatingHistoryParams{
		StockRatingID:   stockRatingHistory.GetStockRatingId(),
		OldCurrentPrice: utils.NumericFromFloat64(stockRatingHistory.GetOldCurrentPrice()),
		NewCurrentPrice: utils.NumericFromFloat64(stockRatingHistory.GetNewCurrentPrice()),
		OldUpside:       utils.NumericFromFloat64(float64(*stockRatingHistory.GetOldUpside())),
		NewUpside:       utils.NumericFromFloat64(float64(stockRatingHistory.GetNewUpside())),
	})
	if err != nil {
		return err
	}
	return nil
}
