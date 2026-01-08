package rest

import (
	"net/http"

	"github.com/jabella1/stock-ratings-service/internal/features/stock/app/interfaces"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/app/query"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type StockRatingController struct {
	getStockRatingByTickerQueryHandler interfaces.GetStockRatingByTickerQueryHandler
}

func CreateStockRatingController(e *echo.Echo, stockRatingService interfaces.GetStockRatingByTickerQueryHandler) *StockRatingController {
	StockRatingController := &StockRatingController{getStockRatingByTickerQueryHandler: stockRatingService}
	e.GET("/api/v1/get-stock-rating-by-ticker/:ticker", StockRatingController.GetStockRatingByTicker)
	e.Use(middleware.Recover())
	return StockRatingController
}

func (c *StockRatingController) GetStockRatingByTicker(ctx echo.Context) error {
	ticker := ctx.Param("ticker")

	if ticker == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ticker parameter is required"})
	}

	getStockRatingByTickerQuery := &query.GetStockRatingByTickerQuery{
		Ticker: ticker,
	}

	result, err := c.getStockRatingByTickerQueryHandler.GetStockRatingByTicker(getStockRatingByTickerQuery)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, result)
}
