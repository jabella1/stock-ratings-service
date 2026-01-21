package domainValidation

import "errors"

func ValidateEmptyString(stringToValidate string, fieldName string) error {
	if stringToValidate == "" {
		return errors.New(fieldName + " is required")
	}
	return nil
}

func NonNegativeOrZeroId(id int64, fieldName string) error {
	if id <= 0 {
		return errors.New(fieldName + " cannot be negative or zero.")
	}
	return nil
}
