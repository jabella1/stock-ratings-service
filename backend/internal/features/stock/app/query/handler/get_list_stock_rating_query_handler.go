package handler

import (
	"github.com/jabella1/stock-ratings-service/internal/features/common/utils"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/app/dto"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/app/query"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/domain/pagination"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/domain/repositories"
)

type GetListStockRatingQueryHandler struct {
	stockRatingRepository repositories.StockRatingRepository
}

func CreateGetListStockRatingQueryHandler(stockRatingRepository repositories.StockRatingRepository) *GetListStockRatingQueryHandler {
	return &GetListStockRatingQueryHandler{
		stockRatingRepository: stockRatingRepository,
	}
}

var orderByMap = map[string]string{
	"symbol":       "ticker",
	"companyName":  "company",
	"broker":       "brokerage",
	"actionType":   "action",
	"ratingFrom":   "rating_from",
	"ratingTo":     "rating_to",
	"targetFrom":   "target_from",
	"targetTo":     "target_to",
	"createdAt":    "created_at",
	"upside":       "upside",
	"currentPrice": "current_price",
}

func (h *GetListStockRatingQueryHandler) GetListStockRating(getListStockRatingQuery *query.GetListStockRatingQuery) (*query.GetListStockRatingResult, error) {
	var orderBy = mapOrderBy(getListStockRatingQuery.OrderBy)
	var orderDirection = utils.MapOrderDirection(getListStockRatingQuery.OrderDirection)
	listStockRatings, err := h.stockRatingRepository.GetListStockRating(getListStockRatingQuery.Search,
		getListStockRatingQuery.PageSize, getListStockRatingQuery.PageNumber, orderBy, orderDirection,
		getListStockRatingQuery.MinUpside, getListStockRatingQuery.MinPrice, getListStockRatingQuery.MaxPrice)
	if err != nil {
		return nil, err
	}

	var resultList []dto.GetListStockRatingResult
	for _, stockRating := range *listStockRatings.Results {
		resultList = append(resultList, dto.GetListStockRatingResult{
			Ticker:       stockRating.GetTicker(),
			Company:      stockRating.GetCompany(),
			Brokerage:    stockRating.GetBrokerage(),
			Action:       stockRating.GetAction(),
			RatingFrom:   stockRating.GetRatingFrom(),
			RatingTo:     stockRating.GetRatingTo(),
			TargetFrom:   stockRating.GetTargetFrom(),
			TargetTo:     stockRating.GetTargetTo(),
			Upside:       stockRating.GetUpside(),
			CurrentPrice: stockRating.GetCurrentPrice(),
		})
	}

	return &query.GetListStockRatingResult{
		Results: &resultList,
		Metadata: pagination.CreatePaginationMetadata(*getListStockRatingQuery.PageNumber,
			*getListStockRatingQuery.PageSize, listStockRatings.TotalRecords, int32(len(*listStockRatings.Results))),
	}, nil
}

func mapOrderBy(input *string) string {
	defaultField := "created_at"
	if input == nil {
		return defaultField
	}
	if value, exists := orderByMap[*input]; exists {
		return value
	}
	return defaultField
}
