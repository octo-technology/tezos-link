package model

// RequestsByDayMetrics count requests made on a given date
type RequestsByDayMetrics struct {
	Year  string
	Month string
	Day   string
	Value int
}

// NewRequestsByDayMetrics returns a new RequestsByDayMetrics
func NewRequestsByDayMetrics(year string, month string, day string, value int) *RequestsByDayMetrics {
	return &RequestsByDayMetrics{
		Year:  year,
		Month: month,
		Day:   day,
		Value: value,
	}
}
