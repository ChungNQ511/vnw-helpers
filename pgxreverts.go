package pgxhelpers

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// RevertPgText reverts a pgtype.Text to a string
// It's useful for converting a pgtype.Text to a string
// @param v any - The value to convert to a string
// @return string - The converted string
func RevertPgText(v any) string {
	switch val := v.(type) {
	case pgtype.Text:
		if val.Valid {
			return val.String
		}
	case *pgtype.Text:
		if val != nil && val.Valid {
			return val.String
		}
	}
	return ""
}

// RevertPgDate reverts a pgtype.Date to a time.Time
// It's useful for converting a pgtype.Date to a time.Time
// @param v any - The value to convert to a time.Time
// @return time.Time - The converted time.Time
func RevertPgDate(v any) time.Time {
	switch val := v.(type) {
	case pgtype.Date:
		if val.Valid {
			return val.Time
		}
	}
	return time.Time{}
}

// RevertPgTimestamp reverts a pgtype.Timestamp to a time.Time
// It's useful for converting a pgtype.Timestamp to a time.Time
// @param v any - The value to convert to a time.Time
// @return time.Time - The converted time.Time
func RevertPgTimestamp(v any) time.Time {
	switch val := v.(type) {
	case pgtype.Timestamp:
		if val.Valid {
			return val.Time
		}
	}
	return time.Time{}
}

// RevertPgTimestamptz reverts a pgtype.Timestamptz to a time.Time
// It's useful for converting a pgtype.Timestamptz to a time.Time
// @param v any - The value to convert to a time.Time
// @return time.Time - The converted time.Time
func RevertPgTimestamptz(v any) time.Time {
	switch val := v.(type) {
	case pgtype.Timestamptz:
		if val.Valid {
			return val.Time
		}
	}
	return time.Time{}
}

// RevertPgBool reverts a pgtype.Bool to a bool
// It's useful for converting a pgtype.Bool to a bool
// @param v any - The value to convert to a bool
// @return bool - The converted bool
func RevertPgBool(v any) bool {
	switch val := v.(type) {
	case pgtype.Bool:
		if val.Valid {
			return val.Bool
		}
	}
	return false
}

// RevertFloatField reverts a pgtype.Float4 or pgtype.Float8 to a float64
// It's useful for converting a pgtype.Float4 or pgtype.Float8 to a float64
// @param v T - The value to convert to a float64
// @return float64 - The converted float64
func RevertFloatField[T pgtype.Float4 | pgtype.Float8](v T) float64 {
	switch val := any(v).(type) {
	case pgtype.Float4:
		if val.Valid {
			return float64(val.Float32)
		}
	case pgtype.Float8:
		if val.Valid {
			return val.Float64
		}
	}
	return 0
}

// RevertIntField reverts a pgtype.Int2 or pgtype.Int4 or pgtype.Int8 to an int
// It's useful for converting a pgtype.Int2 or pgtype.Int4 or pgtype.Int8 to an int
// @param v T - The value to convert to an int
// @return int - The converted int
func RevertIntField[T pgtype.Int2 | pgtype.Int4 | pgtype.Int8](v T) int64 {
	switch val := any(v).(type) {
	case pgtype.Int2:
		if val.Valid {
			return int64(val.Int16)
		}
	case pgtype.Int4:
		if val.Valid {
			return int64(val.Int32)
		}
	case pgtype.Int8:
		if val.Valid {
			return val.Int64
		}
	}
	return 0
}
