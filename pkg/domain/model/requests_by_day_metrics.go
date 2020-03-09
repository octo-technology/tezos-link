package model

type RequestsByDayMetrics struct {
	Year  string
	Month string
	Day   string
	Value int
}

func NewRequestsByDayMetrics(year string, month string, day string, value int) *RequestsByDayMetrics {
	return &RequestsByDayMetrics{
		Year:  year,
		Month: month,
		Day:   day,
		Value: value,
	}
}
