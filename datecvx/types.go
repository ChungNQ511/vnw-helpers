package datecvx

import "time"

// TimeFormat is a type for time formats
// It's useful for formatting and parsing times
// @param t time.Time or *time.Time - The time to format
// @param format TimeFormat - The format of the time
// @return string - The formatted time
type TimeFormat string

const (
	Date_DDMMYYYY        TimeFormat = "02/01/2006"
	Date_YYYYMMDD        TimeFormat = "2006/01/02"
	Date_DDMMYYYY_HHMM   TimeFormat = "02/01/2006 15:04"
	Date_DDMMYYYY_HHMMSS TimeFormat = "02/01/2006 15:04:05"
	Time_HHMMSS          TimeFormat = "15:04:05"
	Time_HHMM            TimeFormat = "15:04"
	DateTime_RFC3339     TimeFormat = time.RFC3339
)
