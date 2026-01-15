package handler

import (
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

func (h *GetListStockRatingQueryHandler) GetListStockRating(getListStockRatingQuery *query.GetListStockRatingQuery) (*query.GetListStockRatingResult, error) {
	listStockRatings, err := h.stockRatingRepository.GetListStockRating(getListStockRatingQuery)
	if err != nil {
		return nil, err
	}

	var resultList []dto.GetListStockRatingResult
	for _, stockRating := range *listStockRatings.Results {
		resultList = append(resultList, dto.GetListStockRatingResult{
			Ticker:     stockRating.GetTicker(),
			Company:    stockRating.GetCompany(),
			Brokerage:  stockRating.GetBrokerage(),
			Action:     stockRating.GetAction(),
			RatingFrom: stockRating.GetRatingFrom(),
			RatingTo:   stockRating.GetRatingTo(),
			TargetFrom: stockRating.GetTargetFrom(),
			TargetTo:   stockRating.GetTargetTo(),
		})
	}

	return &query.GetListStockRatingResult{
		Results: &resultList,
		Metadata: pagination.CreatePaginationMetadata(*getListStockRatingQuery.PageNumber,
			*getListStockRatingQuery.PageSize, listStockRatings.TotalRecords, int32(len(*listStockRatings.Results))),
	}, nil
}
