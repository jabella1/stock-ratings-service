package valueobjects

import "errors"

type Price float64

func CreatePrice(value *float64) (Price, error) {
	if value == nil || *value < 0 {
		return 0, errors.New("price must be positive")
	}
	return Price(*value), nil
}

func (p Price) GetValue() float64 {
	return float64(p)
}
