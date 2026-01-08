package postgres

import (
	"github.com/jackc/pgx/v5/pgtype"
)

func Float64FromNumeric(numeric pgtype.Numeric) *float64 {
	f, _ := numeric.Float64Value()
	return &f.Float64
}
