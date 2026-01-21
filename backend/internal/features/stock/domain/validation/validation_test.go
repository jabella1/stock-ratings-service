package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateEmptyString(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		fieldName string
		wantErr   bool
	}{
		{"error when empty", "", "value", true},
		{"nil when no empty", "juan", "value", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateEmptyString(tt.value, tt.fieldName)
			if (err != nil) != tt.wantErr {
				t.Errorf("expected wanterr=%v and got err=%v", tt.wantErr, err)
			}
		})
	}
}

func TestNonNegativeOrZeroId(t *testing.T) {
	t.Run("returns error when id is zero", func(t *testing.T) {
		err := NonNegativeOrZeroId(0, "idField")
		assert.Error(t, err)
		assert.Equal(t, "idField cannot be negative or zero.", err.Error())
	})

	t.Run("returns error when id is negative", func(t *testing.T) {
		err := NonNegativeOrZeroId(-5, "idField")
		assert.Error(t, err)
		assert.Equal(t, "idField cannot be negative or zero.", err.Error())
	})

	t.Run("returns nil when id is positive", func(t *testing.T) {
		err := NonNegativeOrZeroId(10, "idField")
		assert.NoError(t, err)
	})
}
