package rest

import (
	"net/http"

	"github.com/jabella1/stock-ratings-service/internal/features/common/guard"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/app/interfaces"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/app/query"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/domain/pagination"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/interface/api/rest/restrequest"
	"github.com/labstack/echo/v4"
)

type StockRatingController struct {
	getStockRatingByTickerQueryHandler interfaces.GetStockRatingByTickerQueryHandler
	getListStockRatingQueryHandler     interfaces.GetListStockRatingQueryHandler
}

func CreateStockRatingController(e *echo.Echo,
	getStockRatingByTickerQueryHandler interfaces.GetStockRatingByTickerQueryHandler,
	getListStockRatingQueryHandler interfaces.GetListStockRatingQueryHandler) *StockRatingController {
	StockRatingController := &StockRatingController{getStockRatingByTickerQueryHandler: getStockRatingByTickerQueryHandler,
		getListStockRatingQueryHandler: getListStockRatingQueryHandler}

	e.GET("/api/v1/get-stock-rating-by-ticker/:ticker", StockRatingController.GetStockRatingByTicker)
	e.POST("/api/v1/get-list-stock-rating", StockRatingController.GetListStockRating)
	return StockRatingController
}

func (c *StockRatingController) GetStockRatingByTicker(context echo.Context) error {
	ticker := context.Param("ticker")

	if ticker == "" {
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "ticker parameter is required"})
	}

	getStockRatingByTickerQuery := &query.GetStockRatingByTickerQuery{
		Ticker: ticker,
	}

	result, err := c.getStockRatingByTickerQueryHandler.GetStockRatingByTicker(context.Request().Context(), getStockRatingByTickerQuery)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return context.JSON(http.StatusOK, result)
}

func (c *StockRatingController) GetListStockRating(context echo.Context) error {
	getListStockRatingRequest := &restrequest.GetListStockRatingRequest{}
	if err := context.Bind(getListStockRatingRequest); err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}
	var minUpside float32
	var minPrice, maxPrice float64

	getListStockRatingRequest.PageSize = guard.NonNegativeOrDefaultPtr(getListStockRatingRequest.PageSize, pagination.DefaultPageSize)
	getListStockRatingRequest.PageNumber = guard.NonNegativeOrDefaultPtr(getListStockRatingRequest.PageNumber, pagination.DefaultPageNumber)

	minUpside, err := guard.NonNegative(getListStockRatingRequest.MinUpside)
	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "parameter minUpside Cant'be negative"})
	}

	minPrice, err = guard.NonNegative(getListStockRatingRequest.MinPrice)
	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "parameter minPrice Cant'be negative"})
	}
	maxPrice, err = guard.NonNegative(getListStockRatingRequest.MaxPrice)
	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "parameter maxPrice Cant'be negative"})
	}

	getListStockRatingQuery := &query.GetListStockRatingQuery{
		Search:         getListStockRatingRequest.Search,
		PageSize:       getListStockRatingRequest.PageSize,
		PageNumber:     getListStockRatingRequest.PageNumber,
		OrderBy:        getListStockRatingRequest.OrderBy,
		OrderDirection: getListStockRatingRequest.OrderDirection,
		MinUpside:      minUpside,
		MinPrice:       minPrice,
		MaxPrice:       maxPrice,
	}

	result, err := c.getListStockRatingQueryHandler.GetListStockRating(context.Request().Context(), getListStockRatingQuery)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return context.JSON(http.StatusOK, result)
}
