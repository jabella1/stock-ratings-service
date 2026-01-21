package utils

import "testing"

func TestCalculatePercentageChange(t *testing.T) {
	tests := []struct {
		name string
		to   *float64
		from *float64
		want float64
	}{
		{"to nil", nil, PtrFloat64(10.5), 0},
		{"from nil", PtrFloat64(10.5), nil, 0},
		{"from and to nil", nil, nil, 0},
		{"negative to", PtrFloat64(-11), PtrFloat64(20), 0},
		{"negative from", PtrFloat64(11), PtrFloat64(-20), 0},
		{"positive values", PtrFloat64(11), PtrFloat64(20), -0.45},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculatePercentageChange(tt.from, tt.to)
			if got != tt.want {
				t.Errorf("TestCalculatePercentageChange() got=%v, want=%v", got, tt.want)
			}
		})
	}
}
