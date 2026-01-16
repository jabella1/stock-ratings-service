package dto

type GetStockRatingByTickerResult struct {
	ID         int64
	Ticker     string
	Company    string
	Brokerage  *string
	Action     *string
	RatingFrom *string
	RatingTo   *string
	TargetFrom float64
	TargetTo   float64
}
