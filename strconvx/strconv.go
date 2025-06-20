package strconvx

import (
	"encoding/json"
	"strconv"
	"strings"
)

// ConvertToSlice converts a string to a slice of any type
// It returns a slice of the specified type
// If the input is not a valid slice, it returns an empty slice
// @param input string - The input string to convert to a slice
// @return []T - The converted slice
func ConvertToSlice[T any](input string) []T {
	var result []T

	input = strings.TrimSpace(input)
	if input == "" || input == "{}" || input == "null" {
		return result
	}

	// If the input is {...} → convert to [...]
	if strings.HasPrefix(input, "{") && strings.HasSuffix(input, "}") {
		// Split the elements inside
		trimmed := strings.Trim(input, "{}")
		items := strings.Split(trimmed, ",")

		// Detect the type of string
		_, isString := any(result).([]string)

		var builder strings.Builder
		builder.WriteString("[")

		for i, item := range items {
			item = strings.TrimSpace(item)
			if item == "" {
				continue
			}

			if isString {
				builder.WriteString(strconv.Quote(item)) // add ""
			} else {
				builder.WriteString(item)
			}

			if i < len(items)-1 {
				builder.WriteString(",")
			}
		}
		builder.WriteString("]")

		input = builder.String()
	}

	err := json.Unmarshal([]byte(input), &result)
	if err != nil {
		return result
	}

	return result
}
