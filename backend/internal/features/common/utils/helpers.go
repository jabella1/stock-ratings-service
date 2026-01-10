package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func Float64FromNumeric(numeric pgtype.Numeric) *float64 {
	f, _ := numeric.Float64Value()
	return &f.Float64
}
func TimestamptzFromTime(t time.Time) pgtype.Timestamptz {
	var ts pgtype.Timestamptz
	ts.Scan(t)
	return ts
}

func TimeFromTimestamptz(ts pgtype.Timestamptz) time.Time {
	if ts.Valid {
		return ts.Time
	}
	return time.Time{}
}

func NumericFromFloat64(f float64) pgtype.Numeric {
	var n pgtype.Numeric
	// Convert float64 to string first, then scan
	n.Scan(fmt.Sprintf("%.2f", f))
	return n
}

func TextFromString(s string) pgtype.Text {
	var t pgtype.Text
	t.Scan(s)
	return t
}

func NumericFromString(s string) pgtype.Numeric {
	s = strings.TrimSpace(s)
	if s == "" {
		return pgtype.Numeric{Valid: false}
	}

	clean := strings.ReplaceAll(s, "$", "")
	clean = strings.ReplaceAll(clean, ",", "")
	clean = strings.TrimSpace(clean)

	if clean == "" {
		return pgtype.Numeric{Valid: false}
	}

	var numeric pgtype.Numeric
	if err := numeric.Scan(clean); err != nil {
		return pgtype.Numeric{Valid: false}
	}

	return numeric
}

func CalculateOffset(pageNumber, pageSize int32) int32 {
	if pageNumber <= 1 {
		return 0
	}
	if pageSize <= 0 {
		return 0
	}
	return (pageNumber - 1) * pageSize
}
