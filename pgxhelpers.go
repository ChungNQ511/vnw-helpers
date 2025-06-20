package pgxhelpers

import (
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ChungNQ511/vnw-helpers/funcvx"
	"github.com/jackc/pgx/v5/pgtype"
)

// SetTextField sets a string or []byte to a pgtype.Text
// It returns a pgtype.Text with the string value and a boolean indicating if the value is valid
// If the value is not a string or []byte, it returns a pgtype.Text with an empty string and false
// It's useful for converting a string or []byte to a pgtype.Text
// @param s any - The value to convert to a pgtype.Text
// @return pgtype.Text - The converted pgtype.Text
func SetTextField(v any) pgtype.Text {
	switch val := v.(type) {
	case string:
		return pgtype.Text{
			String: val,
			Valid:  val != "",
		}
	case []byte:
		return pgtype.Text{
			String: string(val),
			Valid:  len(val) > 0,
		}
	case fmt.Stringer:
		s := val.String()
		return pgtype.Text{
			String: s,
			Valid:  s != "",
		}
	default:
		return pgtype.Text{
			String: "",
			Valid:  false,
		}
	}
}

// SetFloatField sets a float32 or float64 or *float64 or pgtype.Numeric to a pgtype.Float4 or pgtype.Float8
// It returns a pgtype.Float4 or pgtype.Float8 with the float32 or float64 or *float64 or pgtype.Numeric value and a boolean indicating if the value is valid
// If the value is not a float32 or float64 or *float64 or pgtype.Numeric, it returns a pgtype.Float4 or pgtype.Float8 with a 0 and false
// It's useful for converting a float32 or float64 or *float64 or pgtype.Numeric to a pgtype.Float4 or pgtype.Float8
// @param v any - The value to convert to a pgtype.Float4 or pgtype.Float8
// @return T - The converted pgtype.Float4 or pgtype.Float8
func SetFloatField[T pgtype.Float4 | pgtype.Float8](v any) T {
	var out T

	switch val := v.(type) {
	case float32:
		out = setFloatToPG[T](float64(val))
	case float64:
		out = setFloatToPG[T](val)
	case *float64:
		if val != nil {
			out = setFloatToPG[T](*val)
		}
	default:
		// fallback: Valid=false
	}

	return out
}

// setFloatToPG sets a float64 to a pgtype.Float4 or pgtype.Float8
// It returns a pgtype.Float4 or pgtype.Float8 with the float64 value and a boolean indicating if the value is valid
// If the value is not a float64, it returns a pgtype.Float4 or pgtype.Float8 with a 0 and false
// It's useful for converting a float64 to a pgtype.Float4 or pgtype.Float8
// @param val float64 - The value to convert to a pgtype.Float4 or pgtype.Float8
// @return T - The converted pgtype.Float4 or pgtype.Float8
func setFloatToPG[T pgtype.Float4 | pgtype.Float8](val float64) T {
	var out T
	switch any(out).(type) {
	case pgtype.Float4:
		out = any(pgtype.Float4{
			Float32: float32(val),
			Valid:   true,
		}).(T)
	case pgtype.Float8:
		out = any(pgtype.Float8{
			Float64: val,
			Valid:   true,
		}).(T)
	}
	return out
}

// SetIntField sets an int or int32 or int64 to a pgtype.Int2 or pgtype.Int4 or pgtype.Int8
// It returns a pgtype.Int2 or pgtype.Int4 or pgtype.Int8 with the int or int32 or int64 value and a boolean indicating if the value is valid
// If the value is not an int or int32 or int64, it returns a pgtype.Int2 or pgtype.Int4 or pgtype.Int8 with a 0 and false
// It's useful for converting an int or int32 or int64 to a pgtype.Int2 or pgtype.Int4 or pgtype.Int8
// @param v any - The value to convert to a pgtype.Int2 or pgtype.Int4 or pgtype.Int8
// @return T - The converted pgtype.Int2 or pgtype.Int4 or pgtype.Int8
func SetIntField[T pgtype.Int2 | pgtype.Int4 | pgtype.Int8](v any) T {
	var out T

	switch val := v.(type) {
	case int:
		out = setIntToPG[T](int64(val))
	case int32:
		out = setIntToPG[T](int64(val))
	case int64:
		out = setIntToPG[T](val)
	default:
		// Trả về zero value với Valid=false
		return out
	}
	return out
}

// setIntToPG sets an int64 to a pgtype.Int2 or pgtype.Int4 or pgtype.Int8
// It returns a pgtype.Int2 or pgtype.Int4 or pgtype.Int8 with the int64 value and a boolean indicating if the value is valid
// If the value is not an int64, it returns a pgtype.Int2 or pgtype.Int4 or pgtype.Int8 with a 0 and false
// It's useful for converting an int64 to a pgtype.Int2 or pgtype.Int4 or pgtype.Int8
// @param val int64 - The value to convert to a pgtype.Int2 or pgtype.Int4 or pgtype.Int8
// @return T - The converted pgtype.Int2 or pgtype.Int4 or pgtype.Int8
func setIntToPG[T pgtype.Int2 | pgtype.Int4 | pgtype.Int8](val int64) T {
	var out T

	switch any(out).(type) {
	case pgtype.Int2:
		out = any(pgtype.Int2{
			Int16: int16(val),
			Valid: true,
		}).(T)
	case pgtype.Int4:
		out = any(pgtype.Int4{
			Int32: int32(val),
			Valid: true,
		}).(T)
	case pgtype.Int8:
		out = any(pgtype.Int8{
			Int64: val,
			Valid: true,
		}).(T)
	}

	return out
}

// SetDateField sets a time.Time or *time.Time or string to a pgtype.Date
// It returns a pgtype.Date with the time.Time or *time.Time or string value and a boolean indicating if the value is valid
// If the value is not a time.Time or *time.Time or string, it returns a pgtype.Date with a zero time and false
// It's useful for converting a time.Time or *time.Time or string to a pgtype.Date
// @param v any - The value to convert to a pgtype.Date
// @return pgtype.Date - The converted pgtype.Date
func SetDateField(v any) pgtype.Date {
	switch val := v.(type) {
	case time.Time:
		return pgtype.Date{Time: val, Valid: true}
	case *time.Time:
		if val != nil {
			return pgtype.Date{Time: *val, Valid: true}
		}
	case string:
		return stringToPgDate(val)
	}
	return pgtype.Date{Valid: false}
}

// SetTimestampField sets a time.Time or *time.Time or string to a pgtype.Timestamp
// It returns a pgtype.Timestamp with the time.Time or *time.Time or string value and a boolean indicating if the value is valid
// If the value is not a time.Time or *time.Time or string, it returns a pgtype.Timestamp with a zero time and false
// It's useful for converting a time.Time or *time.Time or string to a pgtype.Timestamp
// @param v any - The value to convert to a pgtype.Timestamp
// @return pgtype.Timestamp - The converted pgtype.Timestamp
func SetTimestampField(v any) pgtype.Timestamp {
	switch val := v.(type) {
	case time.Time:
		return pgtype.Timestamp{Time: val, Valid: true}
	case *time.Time:
		if val != nil {
			return pgtype.Timestamp{Time: *val, Valid: true}
		}
	case string:
		return stringToPgTimestamp(val)
	}
	return pgtype.Timestamp{Valid: false}
}

// StringToPgDate converts a string to a pgtype.Date
// It returns a pgtype.Date with the string value and a boolean indicating if the value is valid
// If the value is not a string, it returns a pgtype.Date with a zero time and false
// It's useful for converting a string to a pgtype.Date
// @param s string - The value to convert to a pgtype.Date
// @return pgtype.Date - The converted pgtype.Date
func stringToPgDate(s string) pgtype.Date {
	// check null, nil ,...
	if !funcvx.NotNull(s) {
		return pgtype.Date{}
	}

	// accept formats
	formats := []string{
		"2006-01-02",          // ISO format
		"2/1/2006",            // d/m/yyyy
		"2/1/2006 15:04:05",   // d/m/yyyy H:M:S
		"01/02/2006",          // mm/dd/yyyy (nếu có data kiểu Mỹ)
		"2006-01-02 15:04:05", // ISO datetime
		"2006/01/02",          // yyyy/mm/dd
	}

	for _, format := range formats {
		if date, err := time.Parse(format, s); err == nil {
			return pgtype.Date{
				Time:  date,
				Valid: true,
			}
		}
	}

	// fallback if not parsed
	return pgtype.Date{}
}

// stringToPgTimestamp converts a string to a pgtype.Timestamp
// It returns a pgtype.Timestamp with the string value and a boolean indicating if the value is valid
// If the value is not a string, it returns a pgtype.Timestamp with a zero time and false
// It's useful for converting a string to a pgtype.Timestamp
// @param input string - The value to convert to a pgtype.Timestamp
// @return pgtype.Timestamp - The converted pgtype.Timestamp
func stringToPgTimestamp(input string) pgtype.Timestamp {
	if !funcvx.NotNull(input) {
		return pgtype.Timestamp{
			Time:  time.Time{},
			Valid: false,
		}
	}

	input = strings.ReplaceAll(input, "/", "-")

	parts := strings.Split(input, " ")
	if len(parts) < 2 {
		// auto add hour if missing
		parts = append(parts, "00:00:00.000000")
	}

	// auto add hour if missing
	dayParts := strings.Split(parts[0], "-")
	if len(dayParts) != 3 {
		return pgtype.Timestamp{}
	}
	// zero-padding
	for i := 0; i < 2; i++ {
		if len(dayParts[i]) == 1 {
			dayParts[i] = "0" + dayParts[i]
		}
	}
	if len(dayParts[2]) == 2 { // 24 -> 2024
		dayParts[2] = "20" + dayParts[2]
	}
	date := fmt.Sprintf("%s-%s-%s", dayParts[2], dayParts[1], dayParts[0]) // yyyy-mm-dd

	// auto add microsecond if missing
	timePart := parts[1]
	if !strings.Contains(timePart, ".") {
		timePart += ".000000"
	} else {
		// pad microsecond if missing
		tp := strings.Split(timePart, ".")
		for len(tp[1]) < 6 {
			tp[1] += "0"
		}
		timePart = tp[0] + "." + tp[1][:6]
	}

	final := fmt.Sprintf("%s %s", date, timePart)

	// parse
	t, err := time.Parse("2006-01-02 15:04:05.000000", final)
	if err != nil {
		return pgtype.Timestamp{}
	}

	return pgtype.Timestamp{Time: t, Valid: true}
}

// PgBool converts a string to a pgtype.Bool
// @param s: string
// @return pgtype.Bool
func PgBool(s string) pgtype.Bool {
	// s can be 0, 1, true, false, "0", "1", "true", "false", "t", "f", "T", "F", null, "", " ", "  "
	if s == "1" || s == "true" || s == "t" || s == "T" {
		return pgtype.Bool{
			Bool:  true,
			Valid: true,
		}
	}

	if s == "0" || s == "false" || s == "f" || s == "F" {
		return pgtype.Bool{
			Bool:  false,
			Valid: true,
		}
	}

	if s == "null" || s == "" || s == " " || s == "  " {
		return pgtype.Bool{
			Bool:  false,
			Valid: false,
		}
	}

	return pgtype.Bool{
		Bool:  false,
		Valid: true,
	}
}

// SetTimestamptzField sets a time.Time or *time.Time or string to a pgtype.Timestamptz
// It returns a pgtype.Timestamptz with the time.Time or *time.Time or string value and a boolean indicating if the value is valid
// If the value is not a time.Time or *time.Time or string, it returns a pgtype.Timestamptz with a zero time and false
// It's useful for converting a time.Time or *time.Time or string to a pgtype.Timestamptz
// @param v any - The value to convert to a pgtype.Timestamptz
// @return pgtype.Timestamptz - The converted pgtype.Timestamptz
func SetTimestamptzField(v any) pgtype.Timestamptz {
	switch val := v.(type) {
	case time.Time:
		return pgtype.Timestamptz{Time: val.UTC(), Valid: true}
	case *time.Time:
		if val != nil {
			return pgtype.Timestamptz{Time: val.UTC(), Valid: true}
		}
	case string:
		layouts := []string{
			time.RFC3339,          // e.g. "2025-06-17T15:04:05Z"
			"2006-01-02 15:04:05", // e.g. "2025-06-17 14:00:00"
			"2006-01-02",          // fallback: "2025-06-17" → time 00:00:00
		}
		for _, layout := range layouts {
			if t, err := time.Parse(layout, val); err == nil {
				return pgtype.Timestamptz{Time: t.UTC(), Valid: true}
			}
		}
	}
	return pgtype.Timestamptz{Valid: false}
}

func stringToPgTimestamptz(input string) pgtype.Timestamptz {
	// check nil| null | ""
	if !funcvx.NotNull(input) {
		return pgtype.Timestamptz{
			Time:  time.Time{},
			Valid: false,
		}
	}

	input = strings.ReplaceAll(input, "/", "-")

	parts := strings.Split(input, " ")
	if len(parts) < 2 {
		parts = append(parts, "00:00:00.000000")
	}

	dayParts := strings.Split(parts[0], "-")
	if len(dayParts) != 3 {
		return pgtype.Timestamptz{}
	}
	// zero-padding
	for i := 0; i < 2; i++ {
		if len(dayParts[i]) == 1 {
			dayParts[i] = "0" + dayParts[i]
		}
	}
	if len(dayParts[2]) == 2 { // 24 -> 2024
		dayParts[2] = "20" + dayParts[2]
	}
	date := fmt.Sprintf("%s-%s-%s", dayParts[2], dayParts[1], dayParts[0]) // yyyy-mm-dd

	timePart := parts[1]
	if !strings.Contains(timePart, ".") {
		timePart += ".000000"
	} else {
		// pad microsecond if missing
		tp := strings.Split(timePart, ".")
		for len(tp[1]) < 6 {
			tp[1] += "0"
		}
		timePart = tp[0] + "." + tp[1][:6]
	}

	final := fmt.Sprintf("%s %s", date, timePart)

	// parse
	t, err := time.Parse("2006-01-02 15:04:05.000000", final)
	if err != nil {
		return pgtype.Timestamptz{}
	}

	return pgtype.Timestamptz{Time: t, Valid: true}
}

// SetNumericField sets an int or int32 or int64 or float64 or float32 to a pgtype.Numeric
// It returns a pgtype.Numeric with the int or int32 or int64 or float64 or float32 value and a boolean indicating if the value is valid
// If the value is not an int or int32 or int64 or float64 or float32, it returns a pgtype.Numeric with a 0 and false
// It's useful for converting an int or int32 or int64 or float64 or float32 to a pgtype.Numeric
// @param i any - The value to convert to a pgtype.Numeric
// @return pgtype.Numeric - The converted pgtype.Numeric
func SetNumericField(i any) pgtype.Numeric {
	switch res := i.(type) {
	case int:
		return pgtype.Numeric{
			Int:   big.NewInt(int64(res)),
			Exp:   0,
			Valid: true,
		}
	case int32:
		return pgtype.Numeric{
			Int:   big.NewInt(int64(res)),
			Exp:   0,
			Valid: true,
		}
	case int64:
		return pgtype.Numeric{
			Int:   big.NewInt(res),
			Exp:   0,
			Valid: true,
		}
	case float64:
		f := res
		bigF := new(big.Float).SetFloat64(f)
		var num pgtype.Numeric
		num.Int, _ = bigF.Int(nil)
		num.Exp = int32(bigF.MantExp(nil)) - int32(len(bigF.Text('f', -1)))
		num.Valid = true
		return num
	case float32:
		f := res
		bigF := new(big.Float).SetFloat64(float64(f))
		var num pgtype.Numeric
		num.Int, _ = bigF.Int(nil)
		num.Exp = int32(bigF.MantExp(nil)) - int32(len(bigF.Text('f', -1)))
		num.Valid = true
		return num

	default:
		return pgtype.Numeric{
			Exp:   0,
			Valid: false,
		}

	}
}

// SetBoolField sets a bool or int or int32 or int64 or float64 to a pgtype.Bool
// It returns a pgtype.Bool with the bool or int or int32 or int64 or float64 value and a boolean indicating if the value is valid
// If the value is not a bool or int or int32 or int64 or float64, it returns a pgtype.Bool with a false and false
// It's useful for converting a bool or int or int32 or int64 or float64 to a pgtype.Bool
// @param b any - The value to convert to a pgtype.Bool
// @return pgtype.Bool - The converted pgtype.Bool
func SetBoolField(b any) pgtype.Bool {
	switch res := b.(type) {
	case bool:
		return pgtype.Bool{
			Bool:  res,
			Valid: true,
		}
	case int:
		return pgtype.Bool{
			Bool:  res != 0,
			Valid: true,
		}
	case int32:
		return pgtype.Bool{
			Bool:  res != 0,
			Valid: true,
		}
	case int64:
		return pgtype.Bool{
			Bool:  res != 0,
			Valid: true,
		}
	case float64:
		return pgtype.Bool{
			Bool:  res != 0,
			Valid: true,
		}
	case float32:
		return pgtype.Bool{
			Bool:  res != 0,
			Valid: true,
		}
	case string:
		return pgtype.Bool{
			Bool:  res != "0",
			Valid: true,
		}
	default:
		return pgtype.Bool{
			Bool:  false,
			Valid: false,
		}
	}
}
