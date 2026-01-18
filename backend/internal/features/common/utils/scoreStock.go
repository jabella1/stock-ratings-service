package utils

func CalculatePercentageChange(from, to *float64) float64 {
	if from == nil || to == nil || *from == 0 {
		return 0
	}
	return (*to - *from) / *from
}
