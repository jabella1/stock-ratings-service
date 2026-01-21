package handler

import (
	"context"
	"testing"

	"github.com/jabella1/stock-ratings-service/internal/features/common/utils"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/app/query"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/domain/entities"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/domain/pagination"
	"github.com/jabella1/stock-ratings-service/internal/mocks"
)

func TestGetListStockRatingQueryHandler(t *testing.T) {
	context := context.Background()
	orderByCases := map[string]string{
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
		"unknown":      "created_at",
	}

	for inputOrder, expectedOrder := range orderByCases {
		t.Run("orderBy_"+inputOrder, func(t *testing.T) {
			mockStockRatingRepository := mocks.NewStockRatingRepository(t)
			getListStockRatingQueryHandler := CreateGetListStockRatingQueryHandler(context, mockStockRatingRepository)
			pageSize, pageNumber, search, minPrice := int32(5), int32(1), "TestCompany", 100.0
			getListStockRatingQuery := &query.GetListStockRatingQuery{Search: &search, PageSize: &pageSize, PageNumber: &pageNumber, OrderBy: &inputOrder, MinPrice: minPrice}
			stockRatingEntity, err := entities.CreateStockRating("SYM", "Company A", nil, utils.PtrString("target lowered by"), utils.PtrString("buy"), utils.PtrString("hold"), utils.PtrFloat64(10), utils.PtrFloat64(12), 3.4776, 10, 100)

			if err != nil {
				t.Errorf("Error creating Entity")
			}

			mockStockRatingRepository.On("GetListStockRating", context, &search, &pageSize, &pageNumber, expectedOrder, "asc",
				float32(0), float64(100), float64(0),
			).Return(&pagination.PaginatedList[entities.StockRating]{Results: &[]entities.StockRating{*stockRatingEntity}, TotalRecords: 1}, nil)

			result, err := getListStockRatingQueryHandler.GetListStockRating(context, getListStockRatingQuery)
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if result.Metadata.PageNumber != pageNumber || *result.Metadata.PageSize != pageSize || result.Metadata.TotalRecords != 1 {
				t.Errorf("metadata mismatch %+v", result.Metadata)
			}
			if len(*result.Results) != 1 || (*result.Results)[0].Ticker != "SYM" {
				t.Errorf("DTO mapping failed %+v", result.Results)
			}

			mockStockRatingRepository.AssertExpectations(t)
		})
	}
}
