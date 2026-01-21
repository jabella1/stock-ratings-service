package utils

import (
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
)

func TestFloat64FromNumeric(t *testing.T) {

	want := 2.0
	var numeric pgtype.Numeric
	err := numeric.Scan(fmt.Sprint(want))
	if err != nil {
		t.Fatal(err)
	}

	got := Float64FromNumeric(numeric)
	if *got != want {
		t.Errorf("TestFloat64FromNumeric() got=%v, want=%v", *got, want)
	}
}

func TestValidateEmptyString(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		fieldName string
		want      string
		wantErr   bool
	}{
		{"non empty", "juanf", "value", "juanf", false},
		{"empty value", "", "value", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateEmptyString(tt.value, tt.fieldName)
			if (err != nil) != tt.wantErr {
				t.Errorf("TestValidateEmptyString() Error=%v, wantErr=%v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("TestValidateEmptyString() got=%v, want=%v", got, tt.want)
			}
		})
	}
}

func TestNumericFromFloat64(t *testing.T) {
	f := 123.456789
	want := 123.4568

	got := NumericFromFloat64(f)
	val, _ := got.Float64Value()

	if val.Float64 != want {
		t.Errorf("got %v, want %v", val.Float64, want)
	}
}

func TestNumericFromString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		valid bool
		want  float64
	}{
		{"clean", "$1,200.50", true, 1200.50},
		{"spaces", "  50.1  ", true, 50.1},
		{"empty", "", false, 0},
		{"simbols", "$ , ", false, 0},
		{"invalid", "abc", false, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NumericFromString(tt.input)
			if got.Valid != tt.valid {
				t.Errorf("Valid: got %v, want %v", got.Valid, tt.valid)
			}
			if tt.valid {
				f, _ := got.Float64Value()
				if f.Float64 != tt.want {
					t.Errorf("Value: got %v, want %v", f.Float64, tt.want)
				}
			}
		})
	}
}

func TestCalculateOffset(t *testing.T) {
	tests := []struct {
		name string
		page int32
		size int32
		want int32
	}{
		{"pag 1", 1, 10, 0},
		{"pag 2", 2, 10, 10},
		{"pag 3", 3, 10, 20},
		{"invalid pag", 0, 10, 0},
		{"invalid size", 2, 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateOffset(tt.page, tt.size)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapOrderDirection(t *testing.T) {
	desc := "desc"
	asc := "asc"
	other := "random"

	tests := []struct {
		name  string
		input *string
		want  string
	}{
		{"Nil returns asc", nil, "asc"},
		{"Desc returns desc", &desc, "desc"},
		{"Asc returns asc", &asc, "asc"},
		{"another returns asc", &other, "asc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapOrderDirection(tt.input)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTextFromString(t *testing.T) {
	input := "juan f"
	got := TextFromString(input)
	if !got.Valid || got.String != input {
		t.Errorf("Error al convertir texto")
	}
}
