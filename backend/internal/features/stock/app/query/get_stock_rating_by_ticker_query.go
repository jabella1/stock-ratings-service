package query

import "github.com/jabella1/stock-ratings-service/internal/app/features/stock/common"

type GetStockRatingByTickerQuery struct {
	Ticker string
}

type GetStockRatingByTickerResult struct {
	Result *common.StockRatingResult
}
