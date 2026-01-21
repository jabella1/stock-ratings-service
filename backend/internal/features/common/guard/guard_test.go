package guard

import (
	"testing"

	"github.com/jabella1/stock-ratings-service/internal/features/common/utils"
)

func TestNonNegative(t *testing.T) {
	tests := []struct {
		name    string
		value   *float64
		want    float64
		wantErr bool
	}{
		{"nil pointer, return 0", nil, 0, false},
		{"positive value", utils.PtrFloat64(10.5), 10.5, false},
		{"value zero", utils.PtrFloat64(0), 0, false},
		{"negative value - error", utils.PtrFloat64(-1), 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NonNegative(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NonNegative() error =%v, wantError=%v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("NonNegative() got = %v, want = %v", got, tt.want)
			}

		})
	}
}

func TestNonNegativeOrDefaultPtr(t *testing.T) {
	tests := []struct {
		name         string
		value        *int32
		defaultValue int32
		want         int32
		wantErr      bool
	}{
		{"positive value", utils.PtrInt32(2), 5, 2, false},
		{"negative value", utils.PtrInt32(-1), 4, 4, false},
		{"nil value", nil, 5, 5, false},
		{"value zero", utils.PtrInt32(0), 6, 6, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NonNegativeOrDefaultPtr(tt.value, tt.defaultValue)
			if *got != tt.want {
				t.Errorf("TestNonNegativeOrDefaultPtr() got=%v, watn=%v", *got, tt.want)
			}
		})
	}
}
