package guard

func NonNegativeOrDefaultPtr(value *int32, defaultValue int32) *int32 {
	newValue := defaultValue
	if value != nil && *value > 0 {
		newValue = *value
	}
	return &newValue
}
