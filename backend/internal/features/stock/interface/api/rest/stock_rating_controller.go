package rest

import (
	"net/http"

	"github.com/jabella1/stock-ratings-service/internal/features/common/guard"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/app/interfaces"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/app/query"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/domain/pagination"
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

	result, err := c.getStockRatingByTickerQueryHandler.GetStockRatingByTicker(getStockRatingByTickerQuery)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return context.JSON(http.StatusOK, result)
}

func (c *StockRatingController) GetListStockRating(context echo.Context) error {
	getListStockRatingQuery := &query.GetListStockRatingQuery{}
	if err := context.Bind(getListStockRatingQuery); err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}
	getListStockRatingQuery.PageSize = guard.NonNegativeOrDefaultPtr(getListStockRatingQuery.PageSize, pagination.DefaultPageSize)
	getListStockRatingQuery.PageNumber = guard.NonNegativeOrDefaultPtr(getListStockRatingQuery.PageNumber, pagination.DefaultPageNumber)
	result, err := c.getListStockRatingQueryHandler.GetListStockRating(getListStockRatingQuery)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return context.JSON(http.StatusOK, result)
}
