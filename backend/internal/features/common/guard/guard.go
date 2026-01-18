package guard

import "errors"

type Float interface {
	~float32 | ~float64
}

func NonNegativeOrDefaultPtr(value *int32, defaultValue int32) *int32 {
	newValue := defaultValue
	if value != nil && *value > 0 {
		newValue = *value
	}
	return &newValue
}

func NonNegative[T Float](value *T) (T, error) {
	var zero T

	if value == nil {
		return zero, nil
	}

	if *value < 0 {
		return zero, errors.New("value must be non-negative")
	}

	return *value, nil
}
