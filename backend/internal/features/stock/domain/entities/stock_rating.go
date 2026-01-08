package entities

import (
	"errors"
	"time"
)

type StockRating struct {
	id         int64
	ticker     string
	company    string
	brokerage  *string
	action     *string
	ratingFrom *string
	ratingTo   *string
	targetFrom *float64
	targetTo   *float64
	createdAt  time.Time
}

func CreateStockRating(ticker, company string, brokerage, action, ratingFrom,
	ratingTo *string, targetFrom, targetTo *float64) (*StockRating, error) {
	if err := validateEmptyString(ticker, "ticker"); err != nil {
		return nil, err
	}
	if err := validateEmptyString(company, "company"); err != nil {
		return nil, err
	}
	return &StockRating{
		ticker:     ticker,
		company:    company,
		brokerage:  brokerage,
		action:     action,
		ratingFrom: ratingFrom,
		ratingTo:   ratingTo,
		targetFrom: targetFrom,
		targetTo:   targetTo,
		createdAt:  time.Now(),
	}, nil
}

func validateEmptyString(stringToValidate string, fieldName string) error {
	if stringToValidate == "" {
		return errors.New(fieldName + " is required")
	}
	return nil
}

func (sr *StockRating) GetTicker() string {
	return sr.ticker
}

func (sr *StockRating) GetCompany() string {
	return sr.company
}

func (sr *StockRating) GetBrokerage() *string {
	return sr.brokerage
}

func (sr *StockRating) GetAction() *string {
	return sr.action
}

func (sr *StockRating) GetRatingFrom() *string {
	return sr.ratingFrom
}

func (sr *StockRating) GetRatingTo() *string {
	return sr.ratingTo
}

func (sr *StockRating) GetTargetFrom() *float64 {
	return sr.targetFrom
}

func (sr *StockRating) GetTargetTo() *float64 {
	return sr.targetTo
}
