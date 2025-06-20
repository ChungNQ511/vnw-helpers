package datecvx

import "time"

// FormatTime formats a time.Time or *time.Time to a string
// It's useful for formatting a time.Time or *time.Time to a string
// @param t time.Time or *time.Time - The time to format
// @param format TimeFormat - The format of the time
// @return string - The formatted time
func FormatTime[T time.Time | *time.Time](t T, format TimeFormat) string {
	if isZero(t) {
		return ""
	}
	switch v := any(t).(type) {
	case time.Time:
		return v.Format(string(format))
	case *time.Time:
		if v == nil {
			return ""
		}
		return v.Format(string(format))
	default:
		return ""
	}
}

// isZero checks if a time.Time or *time.Time is zero
// It's useful for checking if a time.Time or *time.Time is zero
// @param t time.Time or *time.Time - The time to check
// @return bool - True if the time is zero, false otherwise
func isZero[T time.Time | *time.Time](t T) bool {
	switch v := any(t).(type) {
	case time.Time:
		return v.IsZero()
	case *time.Time:
		return v == nil || v.IsZero()
	default:
		return true
	}
}

// FormatTimeCustom formats a time.Time or *time.Time to a string
// It's useful for formatting a time.Time or *time.Time to a string
// @param v any - The time to format
// @param format TimeFormat - The format of the time
// @return string - The formatted time
func FormatTimeCustom(v any, format TimeFormat) string {
	switch t := v.(type) {
	case time.Time:
		if t.IsZero() {
			return ""
		}
		return t.Format(string(format))
	case *time.Time:
		if t == nil || t.IsZero() {
			return ""
		}
		return t.Format(string(format))
	default:
		return ""
	}
}
