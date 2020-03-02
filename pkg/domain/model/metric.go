package model

import "time"

type Metric struct {
	Request *Request
	Date    time.Time
}

func NewMetric(request *Request, date time.Time) Metric {
	return Metric{
		Request: request,
		Date:    date,
	}
}
