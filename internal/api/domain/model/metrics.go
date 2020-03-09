package model

import "github.com/octo-technology/tezos-link/backend/pkg/domain/model"

// Metrics contains the fields to represent a metrics
type Metrics struct {
	RequestsCount int
	RequestsByDay []*model.RequestsByDayMetrics
}

// NewMetrics returns a new metrics
func NewMetrics(requestsCount int, requestsByDay []*model.RequestsByDayMetrics) Metrics {
	return Metrics{
		RequestsCount: requestsCount,
		RequestsByDay: requestsByDay,
	}
}
