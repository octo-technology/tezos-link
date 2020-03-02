package model

// Metrics contains the fields to represent a metrics
type Metrics struct {
	RequestsCount int
}

// NewMetrics returns a new metrics
func NewMetrics(requestsCount int) Metrics {
	return Metrics{
		RequestsCount: requestsCount,
	}
}
