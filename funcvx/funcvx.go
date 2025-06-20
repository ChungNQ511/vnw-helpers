package funcvx

import (
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func NotNull(input any) bool {
	switch v := input.(type) {
	case string:
		//  s == "null" || s == "" || s == " " || s == "  " / strims
		return v != "null" && v != "NULL" && strings.TrimSpace(v) != "" && v != "nil"
	case int:
		return v != 0
	case int32:
		return v != 0
	case int64:
		return v != 0
	case float64:
		return v != 0
	case bool:
		return v
	case time.Time:
		return !v.IsZero()
	case pgtype.Date:
		return v.Valid
	case pgtype.Timestamp:
		return v.Valid
	case pgtype.Int4:
		return v.Valid
	case pgtype.Int8:
		return v.Valid
	case pgtype.Float8:
		return v.Valid
	case pgtype.Bool:
		return v.Valid
	default:
		return false
	}
}
