package command

import "time"

type SaveStockRatingCommand struct {
	Ticker       string
	Company      string
	Brokerage    *string
	Action       string
	RatingFrom   string
	RatingTo     string
	TargetFrom   float64
	TargetTo     float64
	CurrentPrice float64
}

type SaveStockRatingResult struct {
	Ticker       string    `json:"ticker"`
	Company      string    `json:"company"`
	Brokerage    *string   `json:"brokerage"`
	Action       *string   `json:"action"`
	RatingFrom   *string   `json:"rating_from"`
	RatingTo     *string   `json:"rating_to"`
	TargetFrom   float64   `json:"target_from"`
	TargetTo     float64   `json:"target_to"`
	Upside       float32   `json:"upside"`
	CurrentPrice float64   `json:"currentPrice"`
	CreateAt     time.Time `json:"createAt"`
}
