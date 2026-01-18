package entities

import (
	"errors"
	"time"

	valueobjects "github.com/jabella1/stock-ratings-service/internal/features/stock/domain/valueObjects"
)

type StockRating struct {
	id           int64
	ticker       string
	company      string
	brokerage    *string
	action       *string
	ratingFrom   *string
	ratingTo     *string
	targetFrom   valueobjects.Price
	targetTo     valueobjects.Price
	createdAt    time.Time
	upside       float32
	changeTarget float64
	currentPrice valueobjects.Price
}

func CreateStockRating(ticker, company string, brokerage, action, ratingFrom,
	ratingTo *string, targetFrom, targetTo *float64, upside float32, changeTarget, currentPrice float64) (*StockRating, error) {
	if err := validateEmptyString(ticker, "ticker"); err != nil {
		return nil, err
	}
	if err := validateEmptyString(company, "company"); err != nil {
		return nil, err
	}

	targetFromVO, err := valueobjects.CreatePrice(targetFrom)
	if err != nil {
		return nil, err
	}

	targetToVO, err := valueobjects.CreatePrice(targetTo)
	if err != nil {
		return nil, err
	}

	currenPriceVO, err := valueobjects.CreatePrice(&currentPrice)
	if err != nil {
		return nil, err
	}

	return &StockRating{
		ticker:       ticker,
		company:      company,
		brokerage:    brokerage,
		action:       action,
		ratingFrom:   ratingFrom,
		ratingTo:     ratingTo,
		targetFrom:   targetFromVO,
		targetTo:     targetToVO,
		createdAt:    time.Now(),
		upside:       upside,
		changeTarget: changeTarget,
		currentPrice: currenPriceVO,
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

func (sr *StockRating) GetTargetFrom() float64 {
	return sr.targetFrom.GetValue()
}

func (sr *StockRating) GetTargetTo() float64 {
	return sr.targetTo.GetValue()
}

func (sr *StockRating) GetUpside() float32 {
	return sr.upside
}

func (sr *StockRating) GetCurrentPrice() float64 {
	return sr.currentPrice.GetValue()
}
